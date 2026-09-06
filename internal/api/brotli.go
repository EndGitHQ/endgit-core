package api

import (
	"bufio"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/labstack/echo/v5"
)

const brotliEncoding = "br"

type brotliResponseWriter struct {
	http.ResponseWriter
	writer      *brotli.Writer
	status      int
	wroteHeader bool
	started     bool
	compressed  bool
}

func Brotli() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) (err error) {
			res := c.Response()
			res.Header().Add(echo.HeaderVary, echo.HeaderAcceptEncoding)

			if c.Request().Method == http.MethodHead || !acceptsBrotli(c.Request().Header.Get(echo.HeaderAcceptEncoding)) {
				return next(c)
			}

			brw := &brotliResponseWriter{ResponseWriter: res}
			c.SetResponse(brw)
			defer func() {
				if closeErr := brw.Close(); err == nil && closeErr != nil {
					err = closeErr
				}
				c.SetResponse(res)
			}()

			return next(c)
		}
	}
}

func (w *brotliResponseWriter) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}

	w.status = status
	w.wroteHeader = true
}

func (w *brotliResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	if !w.started {
		w.start()
	}
	if !w.compressed {
		return w.ResponseWriter.Write(b)
	}

	return w.writer.Write(b)
}

func (w *brotliResponseWriter) Flush() {
	if !w.started {
		if !w.wroteHeader {
			w.WriteHeader(http.StatusOK)
		}
		w.start()
	}
	if w.compressed {
		_ = w.writer.Flush()
	}
	_ = http.NewResponseController(w.ResponseWriter).Flush()
}

func (w *brotliResponseWriter) Close() error {
	if !w.started {
		if w.wroteHeader {
			w.ResponseWriter.WriteHeader(w.status)
		}
		return nil
	}
	if !w.compressed {
		return nil
	}

	return w.writer.Close()
}

func (w *brotliResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return http.NewResponseController(w.ResponseWriter).Hijack()
}

func (w *brotliResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *brotliResponseWriter) Push(target string, opts *http.PushOptions) error {
	if pusher, ok := w.ResponseWriter.(http.Pusher); ok {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

func (w *brotliResponseWriter) start() {
	w.started = true
	if !canCompressStatus(w.status) || w.Header().Get(echo.HeaderContentEncoding) != "" {
		w.ResponseWriter.WriteHeader(w.status)
		return
	}

	w.Header().Set(echo.HeaderContentEncoding, brotliEncoding)
	w.Header().Del(echo.HeaderContentLength)
	w.ResponseWriter.WriteHeader(w.status)
	w.writer = brotli.NewWriter(w.ResponseWriter)
	w.compressed = true
}

func canCompressStatus(status int) bool {
	return status >= http.StatusOK && status != http.StatusNoContent && status != http.StatusNotModified
}

func acceptsBrotli(header string) bool {
	for _, part := range strings.Split(header, ",") {
		params := strings.Split(strings.TrimSpace(part), ";")
		if !strings.EqualFold(strings.TrimSpace(params[0]), brotliEncoding) {
			continue
		}

		for _, param := range params[1:] {
			key, value, ok := strings.Cut(strings.TrimSpace(param), "=")
			if !ok || !strings.EqualFold(strings.TrimSpace(key), "q") {
				continue
			}
			quality, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err == nil && quality <= 0 {
				return false
			}
		}

		return true
	}

	return false
}

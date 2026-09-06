package api

import (
	"github.com/EndGitHQ/endgit-core/ent"
	"github.com/labstack/echo/v5"
)

func WithEntClient(client *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			ctx := ent.NewContext(c.Request().Context(), client)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

var serviceStartedAt = time.Now().UTC()

type HealthResponse struct {
	Status        string    `json:"status" example:"ok"`
	StartedAt     time.Time `json:"startedAt"`
	Uptime        string    `json:"uptime" example:"1h2m3s"`
	UptimeSeconds int64     `json:"uptimeSeconds" example:"3723"`
	Timestamp     time.Time `json:"timestamp"`
}

// Health godoc
// @Summary Service health check
// @Description Returns service health status and runtime metadata.
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Health(c *echo.Context) error {
	now := time.Now().UTC()
	uptime := now.Sub(serviceStartedAt)

	return c.JSON(http.StatusOK, HealthResponse{
		Status:        "ok",
		StartedAt:     serviceStartedAt,
		Uptime:        uptime.Truncate(time.Second).String(),
		UptimeSeconds: int64(uptime.Seconds()),
		Timestamp:     now,
	})
}

package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/EndGitHQ/endgit-core/ent"
	"github.com/labstack/echo/v5"
)

var serviceStartedAt = time.Now().UTC()

type HealthResponse struct {
	Status   string `json:"status" example:"ok"`
	Database string `json:"database" example:"ok"`
	Uptime   int64  `json:"uptime" example:"1725571200"`
}

type HealthUnavailableResponse struct {
	Status   string `json:"status" example:"not ok"`
	Database string `json:"database" example:"not ok"`
	Uptime   int64  `json:"uptime" example:"1725571200"`
}

// Health godoc
// @Summary Service health check
// @Description Returns service health status and runtime metadata.
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Failure 503 {object} HealthUnavailableResponse
// @Router /health [get]
func Health(c *echo.Context) error {
	status := "not ok"
	statusCode := http.StatusServiceUnavailable
	databaseStatus := "not ok"

	if client := ent.FromContext(c.Request().Context()); client != nil {
		ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
		defer cancel()

		if _, err := client.User.Query().Exist(ctx); err == nil {
			status = "ok"
			statusCode = http.StatusOK
			databaseStatus = "ok"
		}
	}

	return c.JSON(statusCode, HealthResponse{
		Status:   status,
		Database: databaseStatus,
		Uptime:   serviceStartedAt.Unix(),
	})
}

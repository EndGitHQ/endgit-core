package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// Health godoc
// @Summary Service health check
// @Description Returns service health status.
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func Health(c *echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status: "ok",
	})
}

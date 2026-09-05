package api

import (
	"github.com/EndGitHQ/endgit-core/internal/api/handlers"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/health", handlers.Health)

}

package api

import (
	"github.com/EndGitHQ/endgit-core/internal/api/handlers"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	apiV1.Use(Brotli())

	apiV1.GET("/health", handlers.Health)
	apiV1.GET("/plugins", handlers.ListPlugins)
	apiV1.GET("/plugins/:slug", handlers.GetPluginDetails)
	apiV1.GET("/builds/:id", handlers.GetBuildDetails)

}

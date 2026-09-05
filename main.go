package main

import (
	"net/http"
	"os"

	"log/slog"

	_ "github.com/EndGitHQ/endgit-core/docs"
	"github.com/EndGitHQ/endgit-core/internal/api"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
)

//go:generate go run github.com/swaggo/swag/cmd/swag@latest init --parseInternal -g main.go

// @title EndGit API
// @version 0.1.0
// @description Backend API for EndGit services.
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.GET("/docs", func(c *echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)

	api.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("server listening", "port", port)
	if err := e.Start(":" + port); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}

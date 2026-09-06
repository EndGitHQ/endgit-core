package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"log/slog"

	_ "github.com/EndGitHQ/endgit-core/docs"
	"github.com/EndGitHQ/endgit-core/internal/api"
	"github.com/EndGitHQ/endgit-core/internal/db"
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
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, dbCfg, err := db.OpenFromEnv(ctx)
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			slog.Error("failed to close database connection", "error", closeErr)
		}
	}()

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(api.WithEntClient(client))
	e.GET("/docs", func(c *echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	e.GET("/docs/*", echoSwagger.WrapHandler)

	api.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("database connected",
		"driver", dbCfg.Driver,
		"host", dbCfg.Host,
		"port", dbCfg.Port,
		"name", dbCfg.Name,
		"auto_migrate", dbCfg.AutoMigrate,
	)
	slog.Info("server listening", "port", port)
	if err := e.Start(":" + port); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}

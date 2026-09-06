package db

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/EndGitHQ/endgit-core/ent"
	_ "github.com/lib/pq"
)

const (
	defaultDriver   = "postgres"
	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultUser     = "endgit"
	defaultPassword = "endgit"
	defaultName     = "endgit"
	defaultSSLMode  = "disable"
)

type Config struct {
	Driver      string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	SSLMode     string
	DatabaseURL string
	AutoMigrate bool
}

func ConfigFromEnv() Config {
	port, err := strconv.Atoi(getEnv("DB_PORT", strconv.Itoa(defaultPort)))
	if err != nil {
		port = defaultPort
	}

	return Config{
		Driver:      getEnv("DB_DRIVER", defaultDriver),
		Host:        getEnv("DB_HOST", defaultHost),
		Port:        port,
		User:        getEnv("DB_USER", defaultUser),
		Password:    getEnv("DB_PASSWORD", defaultPassword),
		Name:        getEnv("DB_NAME", defaultName),
		SSLMode:     getEnv("DB_SSLMODE", defaultSSLMode),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		AutoMigrate: parseBool(getEnv("DB_AUTO_MIGRATE", "true")),
	}
}

func (c Config) DSN() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.SSLMode,
	)
}

func OpenEntClient(ctx context.Context, cfg Config) (*ent.Client, error) {
	client, err := ent.Open(cfg.Driver, cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("open ent client: %w", err)
	}

	if cfg.AutoMigrate {
		if err := client.Schema.Create(ctx); err != nil {
			_ = client.Close()
			return nil, fmt.Errorf("auto-migrate schema: %w", err)
		}
	}

	return client, nil
}

func OpenFromEnv(ctx context.Context) (*ent.Client, Config, error) {
	cfg := ConfigFromEnv()
	client, err := OpenEntClient(ctx, cfg)
	if err != nil {
		return nil, Config{}, err
	}
	return client, cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func parseBool(value string) bool {
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return true
	}
	return parsed
}

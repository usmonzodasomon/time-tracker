package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var Logger *slog.Logger

func InitLogger(env string) {
	Logger = GetLogger(env)
}

func GetLogger(env string) *slog.Logger {
	if env == envProd {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
}

package logger

import (
	"log/slog"
	"os"
	"strings"
)

// Init создает JSON-логгер с уровнем, переданным из сервиса
func Init(serviceName string, level string) *slog.Logger {
	var slogLevel slog.Level

	switch strings.ToLower(level) {
	case "debug":
		slogLevel = slog.LevelDebug
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelDebug

		tmpHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slogLevel})
		tmpLogger := slog.New(tmpHandler).With("service", serviceName)
		tmpLogger.Warn("Unknown log level, using default DEBUG", "provided", level)
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLevel,
	})

	logger := slog.New(handler).With("service", serviceName)
	return logger
}

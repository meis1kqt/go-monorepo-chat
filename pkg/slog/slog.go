package slog

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	production = "production"
)



func NewLogger(env string) *slog.Logger {

	var logger *slog.Logger

	switch env {
	case local:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case production:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return logger
}
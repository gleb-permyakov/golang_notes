package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	env := os.Getenv("ENV")
	var levelLogging slog.Level
	switch env {
	case "dev":
		levelLogging = slog.LevelDebug
	case "prod":
		levelLogging = slog.LevelInfo
	}
	TextHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: levelLogging,
		ReplaceAttr: ,
	})
	logger := *slog.New(TextHandler)

	return &logger
}

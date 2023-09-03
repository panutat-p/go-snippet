package log_slog

import (
	"log/slog"
	"testing"
)

// https://pcpratheesh.medium.com/an-overview-of-slog-a-structured-logging-package-for-go-6dec67215b9a
func TestSlog_text_handler(t *testing.T) {
	h := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(h)
	logger.Info(
		"Demo",
		slog.String("app-version", "v0.0.1"),
		slog.Int("release-version", 1),
		slog.Float64("point-value", 1.2),
		slog.Bool("status", true),
		slog.Time("time", time.Now()),
		slog.Group(
			"request",
			slog.String("path", "https://example.com"),
			slog.String("method", "GET"),
		),
	)
}

func TestSlog_json_handler(t *testing.T) {
	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)
	logger.Info(
		"Demo",
		slog.String("app-version", "v0.0.1"),
		slog.Int("release-version", 1),
		slog.Float64("point-value", 1.2),
		slog.Bool("status", true),
		slog.Time("time", time.Now()),
		slog.Group(
			"request",
			slog.String("path", "https://example.com"),
			slog.String("method", "GET"),
		),
	)
}

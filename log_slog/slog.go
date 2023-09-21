package log_slog

import (
	"log/slog"
)

var logger *slog.Logger

func init() {
	h := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{AddSource: true},
	)
	logger = slog.New(h)
}

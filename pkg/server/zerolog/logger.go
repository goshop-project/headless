// Package zerolog provides a zerolog based logger we use in our servers
package zerolog

import (
	"darvaza.org/slog"
	"darvaza.org/slog/handlers/filter"
	"darvaza.org/slog/handlers/zerolog"
)

// NewLogger wraps a zerolog.Logger for slog with a given filter level
func NewLogger(level slog.LogLevel) slog.Logger {
	return filter.New(zerolog.New(&zlog), level)
}

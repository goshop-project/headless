// Package logger provides the logger we use in our servers
package logger

import (
	"github.com/darvaza-proxy/slog"
	"github.com/darvaza-proxy/slog/handlers/filter"
	"github.com/darvaza-proxy/slog/handlers/zerolog"
)

// NewLogger wraps a zerolog.Logger for slog with a given filter level
func NewLogger(level slog.LogLevel) slog.Logger {
	return filter.New(zerolog.New(&zlog), level)
}

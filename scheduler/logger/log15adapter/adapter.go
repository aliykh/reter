// Package log15adapter provides a logger that writes to a github.com/inconshreveable/log15.Logger
// log.
package log15adapter

import (
	"context"
	"github.com/aliykh/reter/scheduler/logger"
)

// Log15Logger interface defines the subset of
// github.com/inconshreveable/log15.Logger that this adapter uses.
type Log15Logger interface {
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Warn(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Crit(msg string, ctx ...interface{})
}

type Logger struct {
	l Log15Logger
}

func NewLogger(l Log15Logger) *Logger {
	return &Logger{l: l}
}

func (l *Logger) Log(ctx context.Context, level logger.LogLevel, msg string, data map[string]interface{}) {
	logArgs := make([]interface{}, 0, len(data))
	for k, v := range data {
		logArgs = append(logArgs, k, v)
	}

	switch level {
	case logger.LogLevelTrace:
		l.l.Debug(msg, append(logArgs, "RETER_LOG_LEVEL", level)...)
	case logger.LogLevelDebug:
		l.l.Debug(msg, logArgs...)
	case logger.LogLevelInfo:
		l.l.Info(msg, logArgs...)
	case logger.LogLevelWarn:
		l.l.Warn(msg, logArgs...)
	case logger.LogLevelError:
		l.l.Error(msg, logArgs...)
	default:
		l.l.Error(msg, append(logArgs, "INVALID_RETER_LOG_LEVEL", level)...)
	}
}

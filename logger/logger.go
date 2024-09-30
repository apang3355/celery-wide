package logger

import (
	"github.com/go-celery/wide"
	"github.com/go-celery/wide/config"
)

type Logger struct {
	logSwitch bool
	logger    celerywide.Logger
}

func New(celeryLog *config.InnerLoggerConfig) (celerywide.Logger, error) {
	return &Logger{
		logSwitch: celeryLog.Switch,
		logger:    celeryLog.Logger,
	}, nil
}

func (l *Logger) Info(ctx celerywide.Context, msg string, fields ...any) {
	if !l.logSwitch {
		return
	}

	l.logger.Info(ctx, msg, fields...)
}

func (l *Logger) Warn(ctx celerywide.Context, msg string, fields ...any) {
	if !l.logSwitch {
		return
	}

	l.logger.Warn(ctx, msg, fields...)
}

func (l *Logger) Error(ctx celerywide.Context, msg string, fields ...any) {
	if !l.logSwitch {
		return
	}

	l.logger.Error(ctx, msg, fields...)
}

func (l *Logger) Debug(ctx celerywide.Context, msg string, fields ...any) {
	if !l.logSwitch {
		return
	}

	l.logger.Debug(ctx, msg, fields...)
}

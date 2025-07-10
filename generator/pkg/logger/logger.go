package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	Key = "logger"
)

type Logger struct {
	logger *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, Key, &Logger{logger: logger})
	return ctx, nil
}

func GetFromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(Key).(*Logger); ok {
		return logger
	}
	return nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	if l.logger != nil {
		l.logger.Info(msg, fields...)
	}
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	if l.logger != nil {
		l.logger.Fatal(msg, fields...)
	}
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l.logger != nil {
		l.logger.Error(msg, fields...)
	}
}

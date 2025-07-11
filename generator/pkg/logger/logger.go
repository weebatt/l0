package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
	"time"

	"go.uber.org/zap"
)

const (
	Key = "logger"
)

type Logger struct {
	logger *zap.Logger
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func New(ctx context.Context) (context.Context, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = customTimeEncoder
	config.EncoderConfig.TimeKey = "time"

	logger, err := config.Build()

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

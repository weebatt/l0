package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Key = "logger"
)

type Logger struct {
	logger *zap.SugaredLogger
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
		return nil, err
	}
	return context.WithValue(ctx, Key, &Logger{logger: logger.Sugar()}), nil
}

func GetFromContext(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

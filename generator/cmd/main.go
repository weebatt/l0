package main

import (
	"context"
	"l0/internal/config"
	"l0/internal/kafka"
	"l0/pkg/logger"
	"os"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	ctx := context.Background()
	ctx, err := logger.New(ctx)
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed to initialize logger", zap.Error(err))
	}

	logger.GetFromContext(ctx).Info("successfully initialized logger")

	// Initialize configuration
	kafkaConfig, err := config.New()
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed to read configuration", zap.Error(err))
	}

	logger.GetFromContext(ctx).Info("successfully read configuration")

	if err := kafka.ProduceOrdersToKafka(ctx, kafkaConfig); err != nil {
		logger.GetFromContext(ctx).Error("Kafka produce error", zap.Error(err))
		os.Exit(1)
	}

	os.Exit(0)
}

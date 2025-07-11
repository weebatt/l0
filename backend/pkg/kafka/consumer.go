package kafka

import (
	"context"
	"encoding/json"
	"l0/internal/config"
	"l0/internal/models"
	"l0/pkg/logger"
	"l0/pkg/metrics"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func StartOrderConsumer(ctx context.Context, cfg *config.Config, handleOrder func(ctx context.Context, order *models.Order) error) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Kafka.Brokers,
		Topic:    cfg.Kafka.Topic,
		GroupID:  cfg.Kafka.GroupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			logger.GetFromContext(ctx).Error("kafka read error", zap.Error(err))
			continue
		}
		metrics.KafkaMessagesConsumed.Inc()
		var order models.Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			logger.GetFromContext(ctx).Error("failed to unmarshal order from kafka", zap.Error(err))
			continue
		}
		logger.GetFromContext(ctx).Info("consumed order from kafka", zap.Any("order_uid", order.OrderUID))
		if err := handleOrder(ctx, &order); err != nil {
			logger.GetFromContext(ctx).Error("failed to handle order", zap.Error(err))
		}
	}
}

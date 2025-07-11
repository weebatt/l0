package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"l0/internal/config"
	"l0/internal/faker"

	"github.com/segmentio/kafka-go"
)

func ProduceOrdersToKafka(ctx context.Context, kafkaConfig *config.Kafka) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaConfig.Broker,
		Topic:    kafkaConfig.Topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	for i := 0; i < kafkaConfig.NumOrders; i++ {
		order := faker.GenerateOrder()
		data, err := json.Marshal(order)
		if err != nil {
			return fmt.Errorf("failed to marshal order: %v", err)
		}
		msg := kafka.Message{
			Key:   []byte(order.OrderUID.String()),
			Value: data,
		}
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			return fmt.Errorf("failed to write message to kafka: %v", err)
		}
		fmt.Printf("Produced order %s to Kafka\n", order.OrderUID)
	}
	return nil
}

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Kafka struct {
	Broker    string `yaml:"kafka_brokers" env:"KAFKA_BROKERS" envDefault:"kafka:9092"`
	Topic     string `yaml:"kafka_topic" env:"KAFKA_TOPIC" envDefault:"orders"`
	GroupID   string `yaml:"kafka_group_id" env:"KAFKA_GROUP_ID" envDefault:"order-consumer"`
	NumOrders int    `yaml:"kafka_num_orders" env:"KAFKA_NUM_ORDERS" envDefault:"10"`
}

func New() (*Kafka, error) {
	var kafka Kafka
	if err := cleanenv.ReadConfig("./config/config.yaml", &kafka); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&kafka); err != nil {
		return nil, err
	}

	return &kafka, nil
}

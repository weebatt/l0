package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Postgres postgres `yaml:"postgres"`
	Kafka    Kafka    `yaml:"kafka"`
}

type Server struct {
	Port            string `yaml:"port" env:"PORT" envDefault:"8080"`
	Host            string `yaml:"host" env:"HOST" envDefault:"localhost"`
	ReadTimeout     int    `yaml:"read_timeout" env:"READ_TIMEOUT" envDefault:"5"`
	WriteTimeout    int    `yaml:"write_timeout" env:"WRITE_TIMEOUT" envDefault:"10"`
	IdleTimeout     int    `yaml:"idle_timeout" env:"IDLE_TIMEOUT" envDefault:"15"`
	ShutdownTimeout int    `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" envDefault:"5"`
}

type postgres struct {
	Port     string `yaml:"port" env:"POSTGRES_PORT" envDefault:"5432"`
	Host     string `yaml:"host" env:"POSTGRES_HOST" envDefault:"localhost"`
	User     string `yaml:"user" env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	DBName   string `yaml:"dbname" env:"POSTGRES_DBNAME" envDefault:"postgres_db"`
	SSLMode  string `yaml:"sslmode" env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

type Kafka struct {
	Brokers []string `yaml:"brokers" env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
	Topic   string   `yaml:"topic" env:"KAFKA_TOPIC" envDefault:"orders"`
	GroupID string   `yaml:"group_id" env:"KAFKA_GROUP_ID" envDefault:"order-consumer"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

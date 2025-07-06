package config

import (
	"net/http"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Postgres postgres `yaml:"postgres"`
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

func NewServer(srv Server) *http.Server {
	server := &http.Server{
		Addr:         ":" + srv.Port,
		ReadTimeout:  time.Duration(srv.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(srv.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(srv.IdleTimeout) * time.Second,
	}

	return server
}

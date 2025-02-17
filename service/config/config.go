package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DbConfig struct {
	Url string `env:"DB_URL"`
	Host string `env:"DB_HOST"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type QueueConfig struct {
	Url string `env:"QUEUE_URL"`
	Host string `env:"QUEUE_HOST"`
	Name string `env:"QUEUE_NAME"`
	User string `env:"QUEUE_USER"`
	Password string `env:"QUEUE_PASSWORD"`
}

func NewDbConfig() (*DbConfig, error) {
	config, err := env.ParseAs[DbConfig](); if err != nil {
		return nil, fmt.Errorf("failed to load db config: %w", err)
	}
	return &config, nil
}

func NewQueueConfig() (*QueueConfig, error) {
	config, err := env.ParseAs[QueueConfig](); if err != nil {
		return nil, fmt.Errorf("failed to load queue config: %w", err)
	}
	return &config, nil
}
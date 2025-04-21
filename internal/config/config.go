package config

import (
	"github.com/caarlos0/env/v11"
	"fmt"
	"time"
)

type Config struct {
	Port string `env:"PORT"`
}

type DatabaseConfig struct {
	Name     string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST"`
	User 	 string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type JWT struct {
	SecretKey              string        `env:"JWT_ACCESS_SECRET"`
	RefreshSecretKey       string        `env:"JWT_REFRESH_SECRET"`
	Issuer                 string        `env:"JWT_ISSUER"`
	AccessExpirationHours  time.Duration `env:"JWT_ACCESS_EXPIRATION"`
	RefreshExpirationHours time.Duration `env:"JWT_REFRESH_EXPIRATION"`
}

func Load() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parsing config: %w", err)
	}
	return cfg, nil
}
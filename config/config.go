package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path/filepath"
)

type Config struct {
	AppName  string `env:"APP_NAME" env-default:"MyApp"`
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string `env:"APP_HOST" env-default:"localhost"`
	Port int    `env:"APP_PORT" env-default:"8000"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	directory, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(directory, ".env")
	err = cleanenv.ReadConfig(envPath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

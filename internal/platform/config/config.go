package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Load() (Config, error) {
	return dbLoad()
}
func dbLoad() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("DB_PORT invalid: %w", err)
	}

	cfg := Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

	if cfg.Database.Host == "" ||
		cfg.Database.User == "" ||
		cfg.Database.Password == "" ||
		cfg.Database.Name == "" {
		return Config{}, fmt.Errorf("incomplete database config")
	}

	if cfg.Database.Port <= 0 || cfg.Database.Port > 65535 {
		return Config{}, fmt.Errorf("DB_PORT out of range")
	}

	return cfg, nil
}

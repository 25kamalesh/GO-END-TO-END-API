package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URI string
	PORT         string
	JWT_SECRET   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		DATABASE_URI: os.Getenv("DATABASE_URI"),
		PORT:         os.Getenv("PORT"),
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
	}
	return cfg, nil
}

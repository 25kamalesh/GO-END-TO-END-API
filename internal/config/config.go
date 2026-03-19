package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URI string
	PORT string
}


func LoadConfig() (*Config , error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		DATABASE_URI: os.Getenv("DATABASE_URI"),
		PORT: os.Getenv("PORT"),
	}
	return cfg, nil
}
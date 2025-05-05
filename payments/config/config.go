package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MercadoPagoToken string
	DB               DBConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		MercadoPagoToken: os.Getenv("MERCADOPAGO_ACCESS_TOKEN"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}

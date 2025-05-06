package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MercadoPagoToken string
	DatabaseURL      string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		MercadoPagoToken: os.Getenv("MERCADOPAGO_ACCESS_TOKEN"),
		DatabaseURL:      os.Getenv("DATABASE_URL"),
	}
}

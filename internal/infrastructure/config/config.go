package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}
}

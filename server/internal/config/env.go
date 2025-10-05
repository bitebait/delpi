package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func IsDevMode() bool {
	return os.Getenv("ENV") == "development"
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Archivo .env no encontrado")
	}
}

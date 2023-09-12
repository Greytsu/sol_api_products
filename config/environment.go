package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file: ", err)
	}

	return &Config{
		Port: os.Getenv("PORT"),
	}
}

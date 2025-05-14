package config

import (
	"log"
	"os"
	"todolist/internal/repository"

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
		Port: os.Getenv("SERVERPORT"),
	}
}

func InitDbConfig() *repository.DBConfig {
	return &repository.DBConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		Username: os.Getenv("DBUSERNAME"),
		Password: os.Getenv("DBPASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	}
}

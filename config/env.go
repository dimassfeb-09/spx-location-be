package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func loadConfig() Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PostgresConfig: PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}
}

func GetConfig() Config {
	return loadConfig()
}

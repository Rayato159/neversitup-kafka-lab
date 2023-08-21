package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App   App
		Kafka Kafka
	}

	App struct {
		Name string
		Url  string
	}

	Kafka struct {
		Url string
	}
)

func NewConfig(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file:", err.Error())
	}
	return Config{
		App: App{
			Name: os.Getenv("APP_NAME"),
			Url:  os.Getenv("APP_URL"),
		},
		Kafka: Kafka{
			Url: os.Getenv("KAFKA_URL"),
		},
	}
}

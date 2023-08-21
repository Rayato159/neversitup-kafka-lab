package main

import (
	"log"
	"os"

	"github.com/Rayato159/neversitup-kafka-lab/src/config"
	"github.com/Rayato159/neversitup-kafka-lab/src/server"
)

func main() {
	// Initialize Config
	cfg := config.NewConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	server.NewServer(&cfg, cfg.App.Name).Start()
}

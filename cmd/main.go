package main

import (
	"avito-shop-service/internal/app"
	"avito-shop-service/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}

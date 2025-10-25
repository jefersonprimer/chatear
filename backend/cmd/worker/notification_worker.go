package main

import (
	"log"

	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/infrastructure"
)

func main() {
	cfg := config.LoadConfig()

	infra, err := infrastructure.NewInfrastructure(cfg.SupabaseConnectionString, cfg.RedisURL, cfg.NatsURL)
	if err != nil {
		log.Fatalf("Error initializing infrastructure: %v", err)
	}
	defer infra.Close()

	log.Println("Starting notification worker...")
	select {}
}

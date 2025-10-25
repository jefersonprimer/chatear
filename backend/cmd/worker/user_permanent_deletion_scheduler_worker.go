package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/infrastructure"
	userApp "github.com/jefersonprimer/chatear/backend/internal/user/application"
	userInfra "github.com/jefersonprimer/chatear/backend/internal/user/infrastructure"
)

func main() {
	cfg := config.LoadConfig()

	infra, err := infrastructure.NewInfrastructure(cfg.SupabaseConnectionString, cfg.RedisURL, "") // NATS not needed for this worker
	if err != nil {
		log.Fatalf("Error initializing infrastructure: %v", err)
	}
	defer infra.Close()

	// Initialize repositories
	userDeletionRepo := userInfra.NewPostgresUserDeletionRepository(infra.DB)
	deletionCapacityRepo := userInfra.NewPostgresDeletionCapacityRepository(infra.DB)

	// Initialize use case
	scheduler := userApp.NewSchedulePermanentDeletions(userDeletionRepo, deletionCapacityRepo, cfg.MaxEmailsPerDay)

	// Run the scheduler periodically
	ticker := time.NewTicker(1 * time.Hour) // Run every hour, adjust as needed
	defer ticker.Stop()

	log.Println("Permanent deletion scheduler worker started.")

	go func() {
		for range ticker.C {
			log.Println("Running permanent deletion scheduler...")
			if err := scheduler.Execute(context.Background()); err != nil {
				log.Printf("Error running permanent deletion scheduler: %v", err)
			}
		}
	}()

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Permanent deletion scheduler worker stopped.")
}

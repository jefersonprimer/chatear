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
	userRepo := userInfra.NewPostgresUserRepository(infra.DB)

	// Initialize use case
	hardDeleter := userApp.NewHardDeleteUsers(userDeletionRepo, userRepo)

	// Run the hard deleter periodically
	ticker := time.NewTicker(24 * time.Hour) // Run once a day, adjust as needed
	defer ticker.Stop()

	log.Println("Hard deletion worker started.")

	go func() {
		for range ticker.C {
			log.Println("Running hard deletion...")
			if err := hardDeleter.Execute(context.Background()); err != nil {
				log.Printf("Error running hard deletion: %v", err)
			}
		}
	}()

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Hard deletion worker stopped.")
}

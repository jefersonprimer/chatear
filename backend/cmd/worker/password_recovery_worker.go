package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/infrastructure"
	notificationApp "github.com/jefersonprimer/chatear/backend/internal/notification/application"
	notificationInfra "github.com/jefersonprimer/chatear/backend/internal/notification/infrastructure"
	notificationWorker "github.com/jefersonprimer/chatear/backend/internal/notification/worker"
	userInfra "github.com/jefersonprimer/chatear/backend/internal/user/infrastructure"
	"github.com/jefersonprimer/chatear/backend/shared/events"
	"github.com/nats-io/nats.go"
)

func main() {
	cfg := config.LoadConfig()

	infra, err := infrastructure.NewInfrastructure("", cfg.RedisURL, cfg.NatsURL)
	if err != nil {
		log.Fatalf("Error initializing infrastructure: %v", err)
	}
	defer infra.Close()

	// Initialize repositories
	notificationRepo := notificationInfra.NewPostgresEmailSendRepository(infra.DB)
	emailLimiter := userInfra.NewRedisEmailLimiter(infra.Redis, cfg)
	oneTimeTokenService := userInfra.NewRedisOneTimeTokenService(infra.Redis, cfg)

	// Initialize notification services
	templateParser := notificationApp.NewHTMLTemplateParser("internal/notification/infrastructure/templates")
	smtpSender := notificationInfra.NewSMTPSender(cfg, templateParser)
	emailSender := notificationApp.NewEmailSender(notificationRepo, smtpSender, emailLimiter)
	emailService := notificationApp.NewEmailService(emailSender, oneTimeTokenService, cfg.AppURL, cfg.MagicLinkExpiry, emailLimiter)

	consumer := notificationWorker.NewPasswordRecoveryConsumer(emailService)

	_, err = infra.NatsConn.Subscribe(events.PasswordRecoveryRequestedSubject, func(msg *nats.Msg) {
		consumer.Consume(context.Background(), msg)
	})
	if err != nil {
		log.Fatalf("Error subscribing to NATS subject: %v", err)
	}

	log.Println("Password recovery worker started. Waiting for events...")

	// Wait for termination signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Password recovery worker stopped.")
}

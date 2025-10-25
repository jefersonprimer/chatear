package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	maxRetries    = 3
	retryInterval = 1 * time.Second
)

// NATSEventBus is a NATS implementation of the domain.EventBus.
type NATSEventBus struct {
	Conn *nats.Conn
}

// NewNATSEventBus creates a new NATSEventBus.
func NewNATSEventBus(conn *nats.Conn) *NATSEventBus {
	return &NATSEventBus{Conn: conn}
}

// Publish publishes an event to the event bus with retry mechanism.
func (b *NATSEventBus) Publish(ctx context.Context, subject string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal event data: %w", err)
	}

	for i := 0; i < maxRetries; i++ {
		err := b.Conn.Publish(subject, jsonData)
		if err == nil {
			return nil
		}
		log.Printf("Attempt %d to publish event to subject %s failed: %v", i+1, subject, err)
		if i < maxRetries-1 {
			time.Sleep(retryInterval)
		}
	}
	return fmt.Errorf("failed to publish event to subject %s after %d attempts", subject, maxRetries)
}

// Subscribe subscribes to a NATS subject.
func (b *NATSEventBus) Subscribe(ctx context.Context, subject string, handler nats.MsgHandler) error {
	_, err := b.Conn.Subscribe(subject, handler)
	if err != nil {
		return fmt.Errorf("failed to subscribe to NATS subject %s: %w", subject, err)
	}
	return nil
}

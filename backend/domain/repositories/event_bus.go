package repositories

import (
	"context"

	"github.com/nats-io/nats.go"
)

// EventBus defines the interface for an event bus.
type EventBus interface {
	Publish(ctx context.Context, subject string, data interface{}) error
	Subscribe(ctx context.Context, subject string, handler nats.MsgHandler) error
}

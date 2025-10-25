package repositories

import (
	"context"
	"time"
)

// BlacklistRepository is an interface for a token blacklist.
type BlacklistRepository interface {
	Add(ctx context.Context, token string, expiration time.Duration) error
	Check(ctx context.Context, token string) (bool, error)
}

package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
)

// NewRedisClient creates a new Redis client
func NewRedisClient(redisURL string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("invalid redis URL: %w", err)
	}
	client := redis.NewClient(opt)
	// Ping to check connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}
	return client, nil
}

// NewNatsClient creates a new NATS client
func NewNatsClient(natsURL string) (*nats.Conn, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}
	return conn, nil
}

// Infrastructure holds all infrastructure components
type Infrastructure struct {
	DB       *pgxpool.Pool
	Redis    *redis.Client
	NatsConn *nats.Conn
}

// NewInfrastructure creates and initializes all infrastructure components
func NewInfrastructure(supabaseConnectionString, redisURL, natsURL string) (*Infrastructure, error) {
	var dbPool *pgxpool.Pool
	var redisClient *redis.Client
	var natsConn *nats.Conn
	var err error

	config, err := pgxpool.ParseConfig(supabaseConnectionString)
	if err != nil {
		log.Printf("Failed to parse database connection string: %v", err)
		dbPool = nil
	} else {


		dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Printf("Failed to create database connection pool: %v", err)
			dbPool = nil
		}
	}

	redisClient, err = NewRedisClient(redisURL)
	if err != nil {
		log.Printf("Failed to create redis client: %v", err)
		redisClient = nil
	}

	if natsURL != "" {
		natsConn, err = NewNatsClient(natsURL)
		if err != nil {
			log.Printf("Failed to create nats client: %v", err)
			natsConn = nil
		}
	}

	return &Infrastructure{
		DB:       dbPool,
		Redis:    redisClient,
		NatsConn: natsConn,
	},
		nil
}

// Close closes all infrastructure connections
func (i *Infrastructure) Close() {
	if i.DB != nil {
		i.DB.Close()
	}
	if i.Redis != nil {
		i.Redis.Close()
	}
	if i.NatsConn != nil {
		i.NatsConn.Close()
	}
}

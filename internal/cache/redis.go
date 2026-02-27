package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new Redis cache client
func NewRedisCache() *RedisCache {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: parseRedisURL(redisURL),
	})

	return &RedisCache{
		client: client,
	}
}

func parseRedisURL(redisURL string) string {
	// Simple parser for redis://host:port format
	// If more complex parsing is needed, use github.com/redis/go-redis option
	if redisURL == "" {
		return "localhost:6379"
	}
	// Remove redis:// prefix if present
	if len(redisURL) > 8 && redisURL[:8] == "redis://" {
		return redisURL[8:]
	}
	return redisURL
}

// Connect establishes a connection to Redis and checks health
func (rc *RedisCache) Connect(ctx context.Context) error {
	pong, err := rc.client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	fmt.Printf("Redis connection successful: %s\n", pong)
	return nil
}

// Close closes the Redis connection
func (rc *RedisCache) Close() error {
	return rc.client.Close()
}

// Get retrieves a value from cache
func (rc *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := rc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key not found")
	}
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// Set stores a value in cache with expiration
func (rc *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return rc.client.Set(ctx, key, data, expiration).Err()
}

// Delete removes a key from cache
func (rc *RedisCache) Delete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return rc.client.Del(ctx, keys...).Err()
}

// Exists checks if a key exists in cache
func (rc *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	result, err := rc.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// GetOrSet retrieves value from cache or computes and stores it
func (rc *RedisCache) GetOrSet(ctx context.Context, key string, dest interface{}, expiration time.Duration, fn func() (interface{}, error)) error {
	// Try to get from cache
	val, err := rc.client.Get(ctx, key).Result()
	if err == nil {
		// Cache hit
		return json.Unmarshal([]byte(val), dest)
	}

	if err != redis.Nil {
		// Some other error, but we'll continue to compute
		fmt.Printf("Cache read error: %v\n", err)
	}

	// Cache miss or error, compute value
	computed, err := fn()
	if err != nil {
		return fmt.Errorf("failed to compute value: %w", err)
	}

	// Store in cache
	if err := rc.Set(ctx, key, computed, expiration); err != nil {
		fmt.Printf("Cache write error: %v\n", err)
		// Don't fail if cache write fails, but log it
	}

	// Unmarshal the computed value
	data, err := json.Marshal(computed)
	if err != nil {
		return fmt.Errorf("failed to marshal computed value: %w", err)
	}

	return json.Unmarshal(data, dest)
}

// FlushAll removes all keys from cache (use with caution)
func (rc *RedisCache) FlushAll(ctx context.Context) error {
	return rc.client.FlushAll(ctx).Err()
}

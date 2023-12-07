package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache represents a simple Redis cache wrapper.
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates and returns a new RedisCache instance.
func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(REDIS_URL), // Replace with your Redis server address
		Password: "",               	// No password by default
		DB:       0,                	// Default DB
	})

	return &RedisCache{client: client}
}

// Set adds a key-value pair to the Redis cache with an expiration time.
func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	if expiration == 0 {
		expiration = 30 * 24 * time.Hour
	}
	return rc.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves the value associated with the given key from the Redis cache.
func (rc *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := rc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s not found in the cache", key)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Delete removes a key from the Redis cache.
func (rc *RedisCache) Delete(key string) error {
	ctx := context.Background()
	return rc.client.Del(ctx, key).Err()
}
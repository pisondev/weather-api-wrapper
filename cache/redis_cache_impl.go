package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheImpl struct {
	Client *redis.Client
}

func (r *RedisCacheImpl) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCacheImpl) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

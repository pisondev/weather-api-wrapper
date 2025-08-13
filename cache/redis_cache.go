package cache

import (
	"context"
	"time"
)

type RedisCache interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	Redis interface {
		Get(ctx context.Context, key string) (res string, err error)
		Set(ctx context.Context, key string, value string) (err error)
		Setex(ctx context.Context, key string, value string, expiration time.Duration) (err error)
		Del(ctx context.Context, key string) (err error)
	}

	redisImpl struct {
		client *redis.Client
	}
)

package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedis() Redis {
	return &redisImpl{
		client: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
		}),
	}
}

func (r *redisImpl) Get(ctx context.Context, key string) (res string, err error) {
	res, err = r.client.Get(ctx, key).Result()
	return
}

func (r *redisImpl) Set(ctx context.Context, key string, value string) (err error) {
	err = r.client.Set(ctx, key, value, 0).Err()
	return
}

func (r *redisImpl) Setex(ctx context.Context, key string, value string, expiration time.Duration) (err error) {
	err = r.client.SetEX(ctx, key, value, expiration).Err()
	return
}

func (r *redisImpl) Del(ctx context.Context, key string) (err error) {
	err = r.client.Del(ctx, key).Err()
	return
}

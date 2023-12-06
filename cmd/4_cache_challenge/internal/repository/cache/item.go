package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/pkg/cache/redis"
)

type itemImpl struct {
	cache redis.Redis
}

func NewItem(
	cache redis.Redis,
) repository.ItemCache {
	return &itemImpl{
		cache: cache,
	}
}

func (r *itemImpl) Get() ([]entity.Item, error) {
	fmt.Println("getting data from cache")
	var (
		result []entity.Item
	)

	items, err := r.cache.Get(context.TODO(), cacheKey)
	if err != nil {
		return nil, err
	}
	if items == "" {
		return nil, errors.New("cache miss")
	}

	err = json.Unmarshal([]byte(items), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *itemImpl) GetByID(id int) (entity.Item, error) {
	// TODO: implement
	return entity.Item{}, nil
}

func (r *itemImpl) Set(items []entity.Item) error {
	val, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return r.cache.Setex(context.TODO(), cacheKey, string(val), 10*time.Second)
}

func (r *itemImpl) SetByID(id int, result entity.Item) error {
	// TODO: implement
	return nil
}

func (r *itemImpl) Del() error {
	// TODO: implement
	return nil
}

func (r *itemImpl) DelByID(id int) error {
	// TODO: implement
	return nil
}

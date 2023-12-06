package proxy

import (
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/repository"
)

type itemProxy struct {
	repository.Item
	cache repository.ItemCache
}

func NewItem(base repository.Item, cache repository.ItemCache) repository.Item {
	return &itemProxy{
		Item:  base,
		cache: cache,
	}
}

func (r *itemProxy) Get() ([]entity.Item, error) {
	var (
		result []entity.Item
		err    error
	)

	result, err = r.cache.Get()
	if err == nil && result != nil {
		return result, nil
	}

	result, err = r.Item.Get()
	if err != nil {
		return result, err
	}

	_ = r.cache.Set(result)

	return result, nil
}

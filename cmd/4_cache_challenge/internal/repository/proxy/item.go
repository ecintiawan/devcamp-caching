package proxy

import (
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository"
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

	go r.cache.Set(result)

	return result, nil
}

func (r *itemProxy) GetByID(id int) (entity.Item, error) {
	var (
		result entity.Item
		err    error
	)

	result, err = r.cache.GetByID(id)
	if err == nil {
		return result, nil
	}

	result, err = r.Item.GetByID(id)
	if err != nil {
		return result, err
	}

	go r.cache.SetByID(id, result)

	return result, nil
}

func (r *itemProxy) Create(model entity.Item) error {
	err := r.Item.Create(model)
	if err != nil {
		return err
	}

	go r.cache.Del()

	return nil
}

func (r *itemProxy) Update(model entity.Item) error {
	err := r.Item.Update(model)
	if err != nil {
		return err
	}

	go func() {
		r.cache.Del()
		r.cache.DelByID(model.ID)
	}()

	return nil
}

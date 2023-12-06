package service

import "github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"

type (
	Item interface {
		GetItems() ([]entity.Item, error)
		GetItem(id int) (entity.Item, error)
		CreateItem(model entity.Item) error
		UpdateItem(model entity.Item) error
	}
)

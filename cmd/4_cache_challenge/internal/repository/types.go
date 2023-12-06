package repository

import "github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"

type (
	Item interface {
		Get() ([]entity.Item, error)
		GetByID(id int) (entity.Item, error)
		Create(model entity.Item) error
		Update(model entity.Item) error
	}
	ItemCache interface {
		Get() ([]entity.Item, error)
		GetByID(id int) (entity.Item, error)
		Set(result []entity.Item) error
		SetByID(id int, result entity.Item) error
		Del() error
		DelByID(id int) error
	}
)

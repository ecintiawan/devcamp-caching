package repository

import "github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/entity"

type (
	Item interface {
		Get() ([]entity.Item, error)
	}
	ItemCache interface {
		Get() ([]entity.Item, error)
		Set(result []entity.Item) error
	}
)

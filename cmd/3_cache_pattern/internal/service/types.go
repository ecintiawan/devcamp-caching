package service

import "github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/entity"

type (
	Item interface {
		GetItems() ([]entity.Item, error)
	}
)

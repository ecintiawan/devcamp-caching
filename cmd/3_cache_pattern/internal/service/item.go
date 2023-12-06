package service

import (
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/repository"
)

type itemImpl struct {
	repo repository.Item
}

func NewItem(
	repo repository.Item,
) Item {
	return &itemImpl{
		repo: repo,
	}
}

func (s *itemImpl) GetItems() ([]entity.Item, error) {
	return s.repo.Get()
}

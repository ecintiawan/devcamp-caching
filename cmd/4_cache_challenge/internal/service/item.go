package service

import (
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository"
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

func (s *itemImpl) GetItem(id int) (entity.Item, error) {
	return s.repo.GetByID(id)
}

func (s *itemImpl) CreateItem(model entity.Item) error {
	return s.repo.Create(model)
}

func (s *itemImpl) UpdateItem(model entity.Item) error {
	return s.repo.Update(model)
}

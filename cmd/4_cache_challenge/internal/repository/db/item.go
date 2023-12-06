package db

import (
	"fmt"

	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/pkg/db"
)

type itemImpl struct {
	db db.DB
}

func NewItem(db db.DB) repository.Item {
	return &itemImpl{
		db: db,
	}
}

func (r *itemImpl) Get() ([]entity.Item, error) {
	fmt.Println("getting data from DB")
	var (
		items []entity.Item
	)

	rows, err := r.db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *itemImpl) GetByID(id int) (entity.Item, error) {
	// TODO: implement
	return entity.Item{}, nil
}

func (r *itemImpl) Create(model entity.Item) error {
	// TODO: implement
	return nil
}

func (r *itemImpl) Update(model entity.Item) error {
	// TODO: implement
	return nil
}

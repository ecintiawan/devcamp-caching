package db

import (
	"fmt"

	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/repository"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/pkg/db"
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

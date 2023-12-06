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
	fmt.Println("getting data from DB")
	var (
		item entity.Item
	)

	if err := r.db.QueryRow("SELECT * FROM items WHERE id = $1 LIMIT 1", id).Scan(&item.ID, &item.Name); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (r *itemImpl) Create(model entity.Item) error {
	query := `
		INSERT INTO items(name)
		VALUES ($1)
	`

	_, err := r.db.Exec(query, model.Name)
	fmt.Println(err)
	return err
}

func (r *itemImpl) Update(model entity.Item) error {
	query := `
		UPDATE items
		SET name = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, model.Name, model.ID)
	fmt.Println(err)
	return err
}

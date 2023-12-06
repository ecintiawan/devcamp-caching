package db

import "database/sql"

type (
	DB interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
	}

	dbImpl struct {
		client *sql.DB
	}
)

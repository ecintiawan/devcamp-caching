package db

import (
	"database/sql"
	"log"
)

func NewDB() DB {
	connStr := "postgres://some_user:some_password@127.0.0.1:5432/some_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	return &dbImpl{
		client: db,
	}
}

func (d *dbImpl) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.client.Query(query, args...)
}

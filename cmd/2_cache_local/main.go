package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

type (
	Item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

const (
	cacheKey = "simulation:item:list"
	cacheDur = 5 * time.Second
)

var (
	db *sql.DB
	c  *cache.Cache
)

func main() {
	var err error

	// init DB
	db, err = initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init Cache
	initCache()

	// new routes
	r := chi.NewRouter()
	r.Get("/items", fetchItemsHandler)
	r.Get("/v2/items", fetchItemsWithCacheHandler)

	// serve http
	port := "9000"
	fmt.Printf("Server is running on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}

func initDB() (*sql.DB, error) {
	connStr := "postgres://some_user:some_password@127.0.0.1:5432/some_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initCache() {
	// TODO: init cache
}

func fetchItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// fetch from DB if cache miss
	items, err := fetchItemsFromDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func fetchItemsWithCacheHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: fetch from cache first

	// fetch from DB if cache miss
	items, err := fetchItemsFromDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// TODO: set data fetched from DB to cache

	json.NewEncoder(w).Encode(items)
}

func fetchItemsFromDB() ([]Item, error) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func fetchItemsFromCache() ([]Item, error) {
	// TODO: implement function

	return nil, nil
}

func setItemsToCache(items []Item) {
	c.Set(cacheKey, items, cacheDur)
}

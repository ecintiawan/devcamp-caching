package main

import (
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	server := app.InitHttp()

	server.InitRoutesAndServe()
}

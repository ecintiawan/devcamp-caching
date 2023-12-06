package main

import (
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	server := app.InitHttp()

	server.InitRoutesAndServe()
}

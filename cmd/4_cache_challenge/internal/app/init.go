package app

import (
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/handler"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository/cache"
	dbRepo "github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository/db"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository/proxy"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/service"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/pkg/cache/redis"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/pkg/db"
)

func InitHttp() *Server {
	db := db.NewDB()
	redis := redis.NewRedis()
	repoItem := dbRepo.NewItem(db)
	cacheItem := cache.NewItem(redis)
	proxyRepoItem := proxy.NewItem(repoItem, cacheItem)
	serviceItem := service.NewItem(proxyRepoItem)
	serverHandler := handler.NewHandler(serviceItem)
	server := NewServer(serverHandler)
	return server
}

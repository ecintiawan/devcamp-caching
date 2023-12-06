package app

import (
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/handler"
	dbRepo "github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/repository/db"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/service"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/pkg/db"
)

func InitHttp() *Server {
	db := db.NewDB()
	repoItem := dbRepo.NewItem(db)
	serviceItem := service.NewItem(repoItem)
	serverHandler := handler.NewHandler(serviceItem)
	server := NewServer(serverHandler)
	return server
}

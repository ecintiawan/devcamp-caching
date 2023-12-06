package app

import (
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/handler"
	dbRepo "github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/repository/db"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/service"
	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/pkg/db"
)

func InitHttp() *Server {
	db := db.NewDB()
	repoItem := dbRepo.NewItem(db)
	serviceItem := service.NewItem(repoItem)
	serverHandler := handler.NewHandler(serviceItem)
	server := NewServer(serverHandler)
	return server
}

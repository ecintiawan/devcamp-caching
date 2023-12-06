package app

import (
	"fmt"
	"net/http"

	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/handler"
	"github.com/go-chi/chi"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(
	handler *handler.Handler,
) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) InitRoutesAndServe() {
	r := chi.NewRouter()
	r.Get("/v2/items", s.handler.GetItems)

	// serve http
	port := "9000"
	fmt.Printf("Server is running on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}

package app

import (
	"fmt"
	"net/http"

	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/handler"
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
	r.Get("/items", s.handler.GetItems)
	r.Get("/items/{id}", s.handler.GetItems)
	r.Post("/items", s.handler.GetItems)
	r.Put("/items/{id}", s.handler.GetItems)

	// serve http
	port := "9000"
	fmt.Printf("Server is running on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}

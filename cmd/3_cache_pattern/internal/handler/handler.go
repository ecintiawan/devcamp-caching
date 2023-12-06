package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ecintiawan/devcamp-caching/cmd/3_cache_pattern/internal/service"
)

type Handler struct {
	serviceItem service.Item
}

func NewHandler(serviceItem service.Item) *Handler {
	return &Handler{
		serviceItem: serviceItem,
	}
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======================== START ========================")
	start := time.Now()
	defer func() {
		latency := time.Since(start)
		fmt.Println("Latency: ", latency)
		fmt.Println("======================== END ========================")
	}()

	w.Header().Set("Content-Type", "application/json")

	// fetch from cache first
	items, err := h.serviceItem.GetItems()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

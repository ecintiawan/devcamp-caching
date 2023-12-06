package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/entity"
	"github.com/ecintiawan/devcamp-caching/cmd/4_cache_challenge/internal/service"
	"github.com/go-chi/chi"
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
	items, err := h.serviceItem.GetItems()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	items, err := h.serviceItem.GetItem(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var (
		model = entity.Item{}
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &model); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = h.serviceItem.CreateItem(model)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Response{
		Message: "SUCCESS",
	})
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var (
		model = entity.Item{}
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &model); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	model.ID, _ = strconv.Atoi(chi.URLParam(r, "id"))

	err = h.serviceItem.UpdateItem(model)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Response{
		Message: "SUCCESS",
	})
}

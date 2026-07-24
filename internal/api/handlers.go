package api

import (
	"net/http"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type Handler struct {
	Service file.Service
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	if h.Service.Ping() {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h Handler) Add(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Add()

	if err != nil {
		
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) Grab(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Grab()

	if err != nil {
		
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Del()

	if err != nil {
		
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) Peek(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Peek()

	if err != nil {
		
	}

	w.WriteHeader(http.StatusOK)
}
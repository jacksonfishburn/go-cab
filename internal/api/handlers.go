package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jacksonfishburn/go-cab/internal/file"
)

type Handler struct {
	Service file.Service
	Token   string
}

func (h Handler) Ping(w http.ResponseWriter, r *http.Request) {
	if h.Service.Ping() {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h Handler) Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	name := r.PathValue("name")
	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}

	record, err := h.Service.Add(name, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) Grab(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	data, err := h.Service.Grab(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	err := h.Service.Del(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) Peek(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.Peek()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

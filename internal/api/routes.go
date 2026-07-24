package api

import (
	"net/http"
)

func (h Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", h.Ping)
	mux.HandleFunc("POST /add", h.Add)
	mux.HandleFunc("GET /grab", h.Grab)
	mux.HandleFunc("DELETE /del", h.Del)
	mux.HandleFunc("GET /peek", h.Peek)
}
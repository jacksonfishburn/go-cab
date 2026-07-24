package api

import (
	"net/http"
)

func (h Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", h.Ping)
	mux.HandleFunc("POST /add/{name}", h.Add)
	mux.HandleFunc("GET /grab/{name}", h.Grab)
	mux.HandleFunc("DELETE /del/{name}", h.Del)
	mux.HandleFunc("GET /peek", h.Peek)
}
package api

import (
	"net/http"
)

func (h Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", h.Ping)
	mux.HandleFunc("/add", h.Add)
	mux.HandleFunc("/grab", h.Grab)
	mux.HandleFunc("/del", h.Del)
	mux.HandleFunc("/peek", h.Peek)
}
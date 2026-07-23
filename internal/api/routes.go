package api

import (
	"net/http"
)

func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", Ping)
	mux.HandleFunc("/add", Add)
	mux.HandleFunc("/grab", Grab)
	mux.HandleFunc("/del", Del)
	mux.HandleFunc("/peek", Peek)
}
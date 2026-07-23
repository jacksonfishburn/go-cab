package api

import (
	"net/http"
)

func Routes(mux *http.ServeMux) {
	api := newAPI()

	mux.HandleFunc("/ping", api.Ping)
	mux.HandleFunc("/add", api.Add)
	mux.HandleFunc("/grab", api.Grab)
	mux.HandleFunc("/del", api.Del)
	mux.HandleFunc("/peek", api.Peek)
}
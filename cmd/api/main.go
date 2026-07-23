package main

import (
	"net/http"
	"github.com/jacksonfishburn/go-cab/internal/api"
)

func main() {
	mux := http.NewServeMux()

	api.Routes(mux)

	http.ListenAndServe("localhost:8080", mux)
}
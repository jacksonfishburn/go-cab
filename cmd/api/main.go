package main

import (
	"net/http"
	"github.com/jacksonfishburn/go-cab/internal/api"
	"github.com/jacksonfishburn/go-cab/internal/file"
	"github.com/jacksonfishburn/go-cab/internal/db"
)

func main() {
	mux := http.NewServeMux()

	var blobstore file.Store
	var mdStore db.MemStore

	service := file.Service{BlobStore: blobstore, MetadataStore: mdStore}
	handler := api.Handler{Service: service}
	handler.Routes(mux)

	http.ListenAndServe("localhost:8080", mux)
}
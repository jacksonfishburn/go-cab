package main

import (
	"net/http"
	"github.com/jacksonfishburn/go-cab/internal/api"
	"github.com/jacksonfishburn/go-cab/internal/file"
	"github.com/jacksonfishburn/go-cab/internal/db"
)

func main() {
	mux := http.NewServeMux()

	blobstore := file.Store{Data: make(map[string][]byte)}
	mdStore := db.MemStore{Records: make(map[string]file.Record)}

	service := file.Service{BlobStore: blobstore, MetadataStore: mdStore}
	handler := api.Handler{Service: service}
	handler.Routes(mux)

	http.ListenAndServe("localhost:8080", mux)
}
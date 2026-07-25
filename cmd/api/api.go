package main

import (
	"net/http"
	"time"
	"log"
	"errors"

	"github.com/jacksonfishburn/go-cab/internal/file"
	"github.com/jacksonfishburn/go-cab/internal/api"
)

type application struct {
	config config
	service file.Service
	handler api.Handler
}

type config struct {
	addr string
	blobstore file.BlobStore
	mdStore file.MetadataStore
}

func (app *application) run() error {
	cfg := app.config

	app.service = file.Service{
		BlobStore: cfg.blobstore,
		MetadataStore: cfg.mdStore,
	}
	app.handler = api.Handler{
		Service: app.service,
	}

	mux := http.NewServeMux()
	app.handler.Routes(mux)

	srv := &http.Server{
		Addr: cfg.addr,
		Handler: mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("server started on %s", cfg.addr)
	err := srv.ListenAndServe()
	log.Println("server closed")

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
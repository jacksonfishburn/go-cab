package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/jacksonfishburn/go-cab/internal/api"
	"github.com/jacksonfishburn/go-cab/internal/file"
)

type application struct {
	config  config
	service file.Service
	handler api.Handler
}

type config struct {
	addr      string
	token     string
	blobstore file.BlobStore
	mdStore   file.MetadataStore
}

func (app *application) run() error {
	cfg := &app.config

	app.service = file.Service{
		BlobStore:     cfg.blobstore,
		MetadataStore: cfg.mdStore,
	}
	app.handler = api.Handler{
		Service: app.service,
		Token:   cfg.token,
	}

	mux := http.NewServeMux()
	app.handler.Routes(mux)

	handler := app.handler.Authorize(mux)

	srv := &http.Server{
		Addr:         cfg.addr,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
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

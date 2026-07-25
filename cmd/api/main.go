package main

import (
	"log"

	"github.com/jacksonfishburn/go-cab/internal/file"
	"github.com/jacksonfishburn/go-cab/internal/db"
	"github.com/jacksonfishburn/go-cab/internal/env"
)

func main() {
	env.Init()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		blobstore: &file.Store{Data: make(map[string][]byte)},
		mdStore: &db.MemStore{Records: make(map[string]file.Record)},
	}
	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
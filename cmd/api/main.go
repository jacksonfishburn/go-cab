package main

import (
	"log"

	"github.com/jacksonfishburn/go-cab/internal/db"
	"github.com/jacksonfishburn/go-cab/internal/env"
	"github.com/jacksonfishburn/go-cab/internal/file"
)

func main() {
	env.Init()

	mdStore, err := db.Open(env.GetString("METADATA_DIR", "go-cab.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer mdStore.Close()

	fileStore, err := file.Open(env.GetString("FILE_STORE_DIR", "data"))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		addr:      env.GetString("ADDR", ":8080"),
		token:     env.GetString("AUTH_TOKEN", ""),
		blobstore: fileStore,
		mdStore:   mdStore,
	}
	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}

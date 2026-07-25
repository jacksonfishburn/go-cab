package main

import (
	"log"

	"github.com/jacksonfishburn/go-cab/internal/db"
	"github.com/jacksonfishburn/go-cab/internal/env"
	"github.com/jacksonfishburn/go-cab/internal/file"
)

func main() {
	env.Init()

	mdStore, err := db.Open(env.GetString("DB_PATH", "go-cab.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer mdStore.Close()


	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		token: env.GetString("AUTH_TOKEN", ""),
		blobstore: &file.Store{Data: make(map[string][]byte)},
		mdStore: mdStore,
	}
	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mhmdhalawi/social/internal/store"
)

func main() {

	// Entry point for the API server
	config := config{
		addr: os.Getenv("ADDR"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: config,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.serve(mux))
}

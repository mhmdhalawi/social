package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// Entry point for the API server
	config := config{
		addr: os.Getenv("ADDR"),
	}
	app := &application{
		config: config,
	}
	mux := app.mount()
	log.Fatal(app.serve(mux))
}

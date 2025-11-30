package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mhmdhalawi/social/internal/db"
	"github.com/mhmdhalawi/social/internal/env"
	"github.com/mhmdhalawi/social/internal/store"
)

func main() {

	// Entry point for the API server
	config := config{
		addr: env.GetString("ADDR"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS"),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME"),
		},
	}

	db, err := db.New(config.db.addr, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Print("database connection opened successfully")

	store := store.NewStorage(db)

	app := &application{
		config: config,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.serve(mux))
}

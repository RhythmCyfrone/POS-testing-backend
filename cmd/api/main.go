package main

import (
	"cyfrone/backend/internal/db"
	"cyfrone/backend/internal/env"
	"cyfrone/backend/internal/store"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: mysql.Config{
			User:   env.GetString("DB_USER", "root"),
			Passwd: env.GetString("DB_PASS", "password"),
			Net:    env.GetString("DB_NET", "tcp"),
			Addr:   env.GetString("DB_ADDR", "localhost:3306"),
			DBName: env.GetString("DB_NAME", "cyfrone"),
		},
	}

	db, err := db.New(cfg.db.FormatDSN())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	store := store.NewSQLStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	err = app.run(mux)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

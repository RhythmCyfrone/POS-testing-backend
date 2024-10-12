package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func New(addr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Println("Connected to database")
	return db, nil
}

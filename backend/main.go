package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zachpmanson/mates/backend/internal/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open DB (sqlite3 used here). This will create a local file `mates.db` in the backend dir.
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "mates.db"
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	// enable foreign keys for sqlite
	_, _ = db.Exec("PRAGMA foreign_keys = ON;")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.AttachHandlers(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

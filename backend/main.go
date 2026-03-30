package backend
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var startTime time.Time

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

type Report struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status: "ok",
		Uptime: time.Since(startTime).String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func reportsHandler(w http.ResponseWriter, r *http.Request) {
	reports := []Report{
		{ID: 1, Title: "Sample report", CreatedAt: time.Now()},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func main() {
	startTime = time.Now()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/reports", reportsHandler)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

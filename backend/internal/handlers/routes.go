package handlers

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func AttachHandlers(db *sql.DB) *http.ServeMux {
	h := &Handler{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.HealthHandler)
	mux.HandleFunc("/api/feeds", h.FeedsHandler)
	mux.HandleFunc("/api/feeds/", h.FeedIdHandler)
	mux.HandleFunc("/api/sightings", h.SightingsHanlder)
	mux.HandleFunc("/api/sightings.rss", h.SightingsGeoRSSHandler)
	mux.HandleFunc("/api/sightings/", h.SightingIdHandler)

	return mux
}

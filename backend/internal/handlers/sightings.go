package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/zachpmanson/mates/backend/internal/sqlc"
	"github.com/zachpmanson/mates/backend/internal/utils"
)

func (h *Handler) SightingsHanlder(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)

	switch r.Method {
	case http.MethodGet:
		// parse query params for ListSightings
		// params: latMin, latMax, longMin, longMax, start, end, feed_id, limit
		qp := r.URL.Query()
		// helper to parse float64
		parseF := func(k string, def float64) float64 {
			if v := qp.Get(k); v != "" {
				if fv, err := strconv.ParseFloat(v, 64); err == nil {
					return fv
				}
			}
			return def
		}
		parseInt64 := func(k string, def int64) int64 {
			if v := qp.Get(k); v != "" {
				if iv, err := strconv.ParseInt(v, 10, 64); err == nil {
					return iv
				}
			}
			return def
		}
		parseTime := func(k string, def time.Time) time.Time {
			if v := qp.Get(k); v != "" {
				if t, err := time.Parse(time.RFC3339, v); err == nil {
					return t
				}
			}
			return def
		}

		latMin := parseF("latMin", -90)
		latMax := parseF("latMax", 90)
		longMin := parseF("longMin", -180)
		longMax := parseF("longMax", 180)
		start := parseTime("start", time.Unix(0, 0))
		end := parseTime("end", time.Now())
		feedID := parseInt64("feed_id", 0)
		limit := parseInt64("limit", 100)

		// Note: the generated ListSightings SQL had repeated/misaligned params; we map accordingly.
		arg := sqlc.ListSightingsParams{
			Lat:         latMin,
			Lat_2:       latMax,
			Long:        longMin,
			Lat_3:       longMax,
			CreatedAt:   start,
			CreatedAt_2: end,
			FeedID:      feedID,
			Limit:       limit,
		}
		items, err := q.ListSightings(r.Context(), arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, items, http.StatusOK)
	case http.MethodPost:
		var in struct {
			CreatedAt string  `json:"created_at"`
			Title     *string `json:"title"`
			Summary   *string `json:"summary"`
			Lat       float64 `json:"lat"`
			Long      float64 `json:"long"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		t := time.Now()
		if in.CreatedAt != "" {
			if tt, err := time.Parse(time.RFC3339, in.CreatedAt); err == nil {
				t = tt
			}
		}
		title := sql.NullString{}
		if in.Title != nil {
			title = sql.NullString{String: *in.Title, Valid: true}
		}
		summary := sql.NullString{}
		if in.Summary != nil {
			summary = sql.NullString{String: *in.Summary, Valid: true}
		}
		created, err := q.CreateSighting(r.Context(), sqlc.CreateSightingParams{CreatedAt: t, Title: title, Summary: summary, Lat: in.Lat, Long: in.Long})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, created, http.StatusCreated)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) SightingIdHandler(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)

	id, err := utils.ParseIDFromPath(r.URL.Path, "/api/sightings/")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		s, err := q.GetSighting(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, s, http.StatusOK)
	case http.MethodPut:
		var in struct {
			CreatedAt string  `json:"created_at"`
			Title     *string `json:"title"`
			Summary   *string `json:"summary"`
			Lat       float64 `json:"lat"`
			Long      float64 `json:"long"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		t := time.Now()
		if in.CreatedAt != "" {
			if tt, err := time.Parse(time.RFC3339, in.CreatedAt); err == nil {
				t = tt
			}
		}
		title := sql.NullString{}
		if in.Title != nil {
			title = sql.NullString{String: *in.Title, Valid: true}
		}
		summary := sql.NullString{}
		if in.Summary != nil {
			summary = sql.NullString{String: *in.Summary, Valid: true}
		}
		if err := q.UpdateSighting(r.Context(), sqlc.UpdateSightingParams{CreatedAt: t, Title: title, Summary: summary, Lat: in.Lat, Long: in.Long, ID: id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if err := q.DeleteSighting(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

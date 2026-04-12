package handlers

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/zachpmanson/mates/backend/internal/sqlc"
	"github.com/zachpmanson/mates/backend/internal/utils"
)

func sightingsGet(q *sqlc.Queries, w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
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
	parseTime := func(k string, def time.Time) string {
		if v := qp.Get(k); v != "" {
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				return t.Format(time.RFC3339)
			}
		}
		return def.Format(time.RFC3339)
	}

	arg := sqlc.ListSightingsParams{
		Lat:         parseF("latMin", -90),
		Lat_2:       parseF("latMax", 90),
		Long:        parseF("longMin", -180),
		Lat_3:       parseF("longMax", 180),
		CreatedAt:   parseTime("start", time.Unix(0, 0)),
		CreatedAt_2: parseTime("end", time.Now()),
		FeedID:      parseInt64("feed_id", 0),
		Limit:       parseInt64("limit", 100),
	}
	items, err := q.ListSightings(r.Context(), arg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, items, http.StatusOK)
}

func sightingsPost(q *sqlc.Queries, w http.ResponseWriter, r *http.Request) {
	var in struct {
		CreatedAt string  `json:"created_at"`
		Title     *string `json:"title"`
		Summary   *string `json:"summary"`
		Lat       float64 `json:"lat"`
		Long      float64 `json:"long"`
		FeedId    int64   `json:"feed_id"`
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
	created, err := q.CreateSighting(r.Context(), sqlc.CreateSightingParams{CreatedAt: t.Format(time.RFC3339), Title: title, Summary: summary, Lat: in.Lat, Long: in.Long, FeedID: in.FeedId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, created, http.StatusCreated)
}

func sightingsGeoRSS(q *sqlc.Queries, w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
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
	parseTime := func(k string, def time.Time) string {
		if v := qp.Get(k); v != "" {
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				return t.Format(time.RFC3339)
			}
		}
		return def.Format(time.RFC3339)
	}

	arg := sqlc.ListSightingsParams{
		Lat:         parseF("latMin", -90),
		Lat_2:       parseF("latMax", 90),
		Long:        parseF("longMin", -180),
		Lat_3:       parseF("longMax", 180),
		CreatedAt:   parseTime("start", time.Unix(0, 0)),
		CreatedAt_2: parseTime("end", time.Now()),
		FeedID:      parseInt64("feed_id", 0),
		Limit:       parseInt64("limit", 100),
	}
	items, err := q.ListSightings(r.Context(), arg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type geoRSSItem struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description,omitempty"`
		PubDate     string   `xml:"pubDate"`
		GeoPoint    string   `xml:"http://www.georss.org/georss point"`
	}
	type geoRSSChannel struct {
		XMLName     xml.Name     `xml:"channel"`
		Title       string       `xml:"title"`
		Description string       `xml:"description"`
		Items       []geoRSSItem `xml:"item"`
	}
	type geoRSSFeed struct {
		XMLName xml.Name      `xml:"rss"`
		Version string        `xml:"version,attr"`
		GeoRSS  string        `xml:"xmlns:georss,attr"`
		Channel geoRSSChannel `xml:"channel"`
	}

	feed := geoRSSFeed{
		Version: "2.0",
		GeoRSS:  "http://www.georss.org/georss",
		Channel: geoRSSChannel{
			Title:       "Mates Sightings",
			Description: "Wildlife sightings feed",
		},
	}
	for _, s := range items {
		item := geoRSSItem{
			Title:    fmt.Sprintf("Sighting #%d", s.ID),
			PubDate:  s.CreatedAt,
			GeoPoint: fmt.Sprintf("%g %g", s.Lat, s.Long),
		}
		if s.Title.Valid {
			item.Title = s.Title.String
		}
		if s.Summary.Valid {
			item.Description = s.Summary.String
		}
		feed.Channel.Items = append(feed.Channel.Items, item)
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml.Header))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	enc.Encode(feed)
}

func (h *Handler) SightingsGeoRSSHandler(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sightingsGeoRSS(q, w, r)
}

func (h *Handler) SightingsHanlder(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)

	switch r.Method {
	case http.MethodGet:
		sightingsGet(q, w, r)
	case http.MethodPost:
		sightingsPost(q, w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func sightingIdGet(q *sqlc.Queries, w http.ResponseWriter, r *http.Request, id int64) {
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
}

func sightingIdPut(q *sqlc.Queries, w http.ResponseWriter, r *http.Request, id int64) {
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
	if err := q.UpdateSighting(r.Context(), sqlc.UpdateSightingParams{CreatedAt: t.Format(time.RFC3339), Title: title, Summary: summary, Lat: in.Lat, Long: in.Long, ID: id}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func sightingIdDelete(q *sqlc.Queries, w http.ResponseWriter, r *http.Request, id int64) {
	if err := q.DeleteSighting(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
		sightingIdGet(q, w, r, id)
	case http.MethodPut:
		sightingIdPut(q, w, r, id)
	case http.MethodDelete:
		sightingIdDelete(q, w, r, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

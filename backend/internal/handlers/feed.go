package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/zachpmanson/mates/backend/internal/sqlc"
	"github.com/zachpmanson/mates/backend/internal/utils"
)

func (h *Handler) FeedsHandler(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)

	switch r.Method {
	case http.MethodGet:
		// list
		feeds, err := q.ListAllFeeds(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, feeds, http.StatusOK)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) FeedIdHandler(w http.ResponseWriter, r *http.Request) {
	q := sqlc.New(h.DB)

	id, err := utils.ParseIDFromPath(r.URL.Path, "/api/feeds/")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		feed, err := q.GetFeedByID(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, feed, http.StatusOK)
	case http.MethodPut:
		var in struct {
			Name string  `json:"name"`
			Desc *string `json:"desc"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		desc := sql.NullString{}
		if in.Desc != nil {
			desc = sql.NullString{String: *in.Desc, Valid: true}
		}
		if err := q.UpdateFeed(r.Context(), sqlc.UpdateFeedParams{Name: in.Name, Desc: desc, ID: id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if err := q.DeleteFeed(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

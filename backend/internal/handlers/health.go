package handlers

import (
	"net/http"
	"time"

	"github.com/zachpmanson/mates/backend/internal/utils"
)

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

var startTime time.Time

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status: "ok",
		Uptime: time.Since(startTime).String(),
	}
	utils.WriteJSON(w, resp, http.StatusOK)
}

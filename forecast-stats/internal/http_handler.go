package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPHandler struct {
	statsService StatsService
}

func NewHTTPHandler(statsService StatsService) *HTTPHandler {
	return &HTTPHandler{
		statsService: statsService,
	}
}

func (h *HTTPHandler) HandleStats(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	ctx := context.Background()

	stats, err := h.statsService.GetStats(ctx, slug)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch stats"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(stats)
}

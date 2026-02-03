package forecast

import (
	"context"
	"encoding/json"
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

func (h *HTTPHandler) HandleWeeklyForecast(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	ctx := context.Background()

	forecast, err := h.statsService.GetForecast(ctx, slug)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch forecast"})
		return
	}

	_ = json.NewEncoder(w).Encode(forecast)
}

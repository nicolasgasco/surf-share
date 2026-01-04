package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"surf-share/app/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BreaksHandler struct {
	pool *pgxpool.Pool
}

func NewBreaksHandler(pool *pgxpool.Pool) *BreaksHandler {
	return &BreaksHandler{pool: pool}
}

func (h *BreaksHandler) HandleBreaks(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.pool.Query(context.Background(), "SELECT id, name, slug FROM app.breaks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var breaks []models.BreakSummary
	for rows.Next() {
		var brk models.BreakSummary
		if err := rows.Scan(&brk.ID, &brk.Name, &brk.Slug); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		breaks = append(breaks, brk)
	}

	resp := struct {
		Count  int                   `json:"count"`
		Breaks []models.BreakSummary `json:"breaks"`
	}{
		Count:  len(breaks),
		Breaks: breaks,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

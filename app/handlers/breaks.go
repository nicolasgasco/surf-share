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

func (h *BreaksHandler) HandleBreakBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	var brk models.Break
	err := h.pool.QueryRow(context.Background(),
		"SELECT id, name, slug, description, coordinates, country, region, city, created_at, updated_at FROM app.breaks WHERE slug=$1", slug).Scan(
		&brk.ID, &brk.Name, &brk.Slug, &brk.Description, &brk.Coordinates,
		&brk.Country, &brk.Region, &brk.City, &brk.CreatedAt, &brk.UpdatedAt)

	if err != nil {
		http.Error(w, "Break not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(brk); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

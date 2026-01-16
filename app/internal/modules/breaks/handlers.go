package breaks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"surf-share/app/internal/adapters"
	"surf-share/app/internal/models"
)

type BreaksHandler struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewBreaksHandler(dbAdapter *adapters.DatabaseAdapter) *BreaksHandler {
	return &BreaksHandler{dbAdapter: dbAdapter}
}

func (h *BreaksHandler) HandleBreaks(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	var breaks []BreaksResponse = make([]BreaksResponse, 0)
	if err := h.dbAdapter.FindMany(ctx, &breaks, "SELECT id, name, slug FROM app.breaks ORDER BY name ASC"); err != nil {
		http.Error(w, "Failed to fetch breaks", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Count  int              `json:"count"`
		Breaks []BreaksResponse `json:"breaks"`
	}{
		Count:  len(breaks),
		Breaks: breaks,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		errorResponse := models.ErrorResponse{}
		errorResponse.Message = "Something went wrong"

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(errorResponse)
	}
}

func (h *BreaksHandler) HandleBreakBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	ctx := context.Background()

	var brk BreakResponse
	err := h.dbAdapter.FindOne(
		ctx,
		&brk,
		`SELECT b.id, b.name, b.slug, b.description, b.coordinates, b.country, b.region, b.city, m.video_url, m.image_urls
		 FROM app.breaks b
		 LEFT JOIN app.breaks_media m ON b.slug = m.break_slug
		 WHERE b.slug = $1`,
		slug,
	)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errorResponse := models.ErrorResponse{}
		errorResponse.Message = fmt.Sprintf("Break with slug '%s' not found", slug)

		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(errorResponse)
		return
	}

	_ = json.NewEncoder(w).Encode(brk)
}

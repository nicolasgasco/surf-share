package breaks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type handlerBreaksResponse struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
}

type handlerBreakResponse struct {
	ID          pgtype.UUID  `json:"id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Description pgtype.Text  `json:"description"`
	Coordinates pgtype.Point `json:"coordinates"`
	Country     string       `json:"country"`
	Region      pgtype.Text  `json:"region"`
	City        pgtype.Text  `json:"city"`
	VideoUrl    pgtype.Text  `json:"video_url"`
	ImageUrls   []string     `json:"image_urls"`
}

type HTTPHandler struct {
	breaksService BreaksService
}

func NewHTTPHandler(breaksService BreaksService) *HTTPHandler {
	return &HTTPHandler{
		breaksService: breaksService,
	}
}

func (h *HTTPHandler) HandleBreaks(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	breaks, err := h.breaksService.GetBreaks(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch breaks", http.StatusInternalServerError)
		return
	}

	handlerBreaks := make([]handlerBreaksResponse, len(breaks))
	for i, b := range breaks {
		handlerBreaks[i] = handlerBreaksResponse{
			ID:   b.ID,
			Name: b.Name,
			Slug: b.Slug,
		}
	}

	resp := struct {
		Count  int                     `json:"count"`
		Breaks []handlerBreaksResponse `json:"breaks"`
	}{
		Count:  len(handlerBreaks),
		Breaks: handlerBreaks,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong"})
	}
}

func (h *HTTPHandler) HandleBreakBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	ctx := context.Background()

	brk, err := h.breaksService.GetBreakBySlug(ctx, slug)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Break with slug '%s' not found", slug),
		})
		return
	}

	handlerBreak := handlerBreakResponse{
		ID:          brk.ID,
		Name:        brk.Name,
		Slug:        brk.Slug,
		Description: brk.Description,
		Coordinates: brk.Coordinates,
		Country:     brk.Country,
		Region:      brk.Region,
		City:        brk.City,
		VideoUrl:    brk.VideoUrl,
		ImageUrls:   brk.ImageUrls,
	}

	_ = json.NewEncoder(w).Encode(handlerBreak)
}

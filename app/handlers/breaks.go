package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"surf-share/app/models"

	"github.com/jackc/pgx/v5"
)

type BreaksHandler struct {
	conn *pgx.Conn
}

func NewBreaksHandler(conn *pgx.Conn) *BreaksHandler {
	return &BreaksHandler{conn: conn}
}

func (h *BreaksHandler) HandleBreaks(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.conn.Query(context.Background(), "SELECT id, name, slug FROM app.breaks")
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

package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleBreaks(w http.ResponseWriter, _ *http.Request) {
	breaks := []string{"La Arena", "Sopelana"}

	resp := struct {
		Count  int      `json:"count"`
		Breaks []string `json:"breaks"`
	}{
		Count:  len(breaks),
		Breaks: breaks,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

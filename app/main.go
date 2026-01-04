package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &homeHandler{})

	// Breaks
	mux.Handle("/breaks", &BreaksHandler{})

	port := os.Getenv("PORT")
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))

}

type BreaksHandler struct{}

func (h *BreaksHandler) ListBreaks(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *BreaksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		h.ListBreaks(w, r)
	}
}

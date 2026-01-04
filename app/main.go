package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)

	// Breaks
	mux.HandleFunc("GET /breaks", handleBreaks)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	filePath := "static/app.html"
	http.ServeFile(w, r, filePath)
}

func handleBreaks(w http.ResponseWriter, r *http.Request) {
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

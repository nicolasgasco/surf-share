package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"forecast-stats/internal"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	tmpl, err := template.ParseFiles("templates/root.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	openMeteoClient := internal.NewOpenMeteoClient()
	breaksClient := internal.NewBreaksClient()
	statsService := internal.NewStatsService(openMeteoClient, breaksClient)
	statsHandler := internal.NewHTTPHandler(statsService)
	mux.HandleFunc("GET /forecast/stats/{slug}", statsHandler.HandleStats)

	log.Printf("Starting forecast-stats service on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

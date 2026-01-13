package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"surf-share/app/internal/adapters"
	"surf-share/app/internal/handlers"
	"surf-share/app/internal/middleware"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	ctx := context.Background()

	dbAdapter := adapters.DatabaseAdapter{}
	if err := dbAdapter.Connect(ctx, connStr); err != nil {
		panic(err)
	}
	defer dbAdapter.Close()

	mux := http.NewServeMux()

	// Breaks
	mux.HandleFunc("GET /breaks", handlers.NewBreaksHandler(&dbAdapter).HandleBreaks)
	mux.HandleFunc("GET /breaks/{slug}", handlers.NewBreaksHandler(&dbAdapter).HandleBreakBySlug)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	port := os.Getenv("PORT")

	// Wrap mux with CORS middleware
	handler := middleware.CORS(mux)

	fmt.Printf("Server is listening to port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("Error starting application")
	}
}

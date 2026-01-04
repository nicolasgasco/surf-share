package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"surf-share/app/handlers"
	"surf-share/app/middleware"
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

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	mux := http.NewServeMux()

	// Breaks
	mux.HandleFunc("GET /breaks", handlers.NewBreaksHandler(pool).HandleBreaks)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)

	// Wrap mux with CORS middleware
	handler := middleware.CORS(mux)
	http.ListenAndServe(":"+port, handler)
}

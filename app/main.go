package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"surf-share/app/handlers"
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

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	mux := http.NewServeMux()

	// Breaks
	mux.HandleFunc("GET /breaks", handlers.NewBreaksHandler(conn).HandleBreaks)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}

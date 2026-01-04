package main

import (
	"fmt"
	"net/http"
	"os"

	"surf-share/app/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Breaks
	mux.HandleFunc("GET /breaks", handlers.HandleBreaks)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	http.ListenAndServe(":"+port, mux)
}

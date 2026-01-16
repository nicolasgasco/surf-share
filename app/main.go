package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"surf-share/app/config"
	"surf-share/app/internal/adapters"
	"surf-share/app/internal/handlers"
	"surf-share/app/internal/middleware"
	"surf-share/app/internal/modules/auth"
	"surf-share/app/internal/modules/breaks"
)

func main() {
	ctx := context.Background()

	connStr, err := config.GetDatabaseConnectionString()
	if err != nil {
		panic(err)
	}

	dbAdapter := adapters.DatabaseAdapter{}
	if err := dbAdapter.Connect(ctx, connStr); err != nil {
		panic(err)
	}
	defer dbAdapter.Close()

	mux := http.NewServeMux()

	breaksModule := breaks.NewBreaksModule(&dbAdapter)
	breaksModule.Register(mux)

	authModule := auth.NewAuthModule(&dbAdapter)
	authModule.Register(mux)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("Error starting application")
	}
}

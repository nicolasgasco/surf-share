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

	breaks.NewBreaksModule(&dbAdapter).Register(mux)

	userRepository := auth.NewRepository(&dbAdapter)
	passwordHasher := auth.NewBcryptHasher()
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenGenerator := auth.NewJWTGenerator(jwtSecret)
	authService := auth.NewAuthService(userRepository, passwordHasher, tokenGenerator)
	httpHandler := auth.NewHTTPHandler(authService)
	mux.HandleFunc("POST /auth/register", httpHandler.HandleRegister)
	mux.HandleFunc("POST /auth/login", httpHandler.HandleLogin)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("Error starting application")
	}
}

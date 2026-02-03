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
	"surf-share/app/internal/modules/forecast"
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

	// Breaks
	breaksRepository := breaks.NewRepository(&dbAdapter)
	breaksService := breaks.NewBreaksService(breaksRepository)
	breaksHandler := breaks.NewHTTPHandler(breaksService)
	mux.HandleFunc("GET /breaks", breaksHandler.HandleBreaks)
	mux.HandleFunc("GET /breaks/{slug}", breaksHandler.HandleBreakBySlug)

	// Auth
	authRepository := auth.NewRepository(&dbAdapter)
	passwordHasher := auth.NewBcryptHasher()
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenGenerator := auth.NewJWTGenerator(jwtSecret)
	authService := auth.NewAuthService(authRepository, passwordHasher, tokenGenerator)
	authHttpHandler := auth.NewHTTPHandler(authService)
	mux.HandleFunc("POST /auth/register", authHttpHandler.HandleRegister)
	mux.HandleFunc("POST /auth/login", authHttpHandler.HandleLogin)

	// Forecast
	forecastRepository := forecast.NewRepository(&dbAdapter)
	openMeteoClient := forecast.NewOpenMeteoClient()
	forecastService := forecast.NewStatsService(forecastRepository, openMeteoClient)
	forecastHandler := forecast.NewHTTPHandler(forecastService)
	mux.HandleFunc("GET /forecast/{slug}", forecastHandler.HandleWeeklyForecast)

	mux.HandleFunc("GET /", handlers.HandleRoot)

	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	fmt.Printf("Server is listening to port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println("Error starting application")
	}
}

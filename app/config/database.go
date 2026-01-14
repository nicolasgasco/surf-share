package config

import (
	"fmt"
	"os"
)

func GetDatabaseConnectionString() (string, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return "", fmt.Errorf("DB_HOST environment variable is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return "", fmt.Errorf("DB_PORT environment variable is not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return "", fmt.Errorf("DB_USER environment variable is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return "", fmt.Errorf("DB_PASSWORD environment variable is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return "", fmt.Errorf("DB_NAME environment variable is not set")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	return connStr, nil
}

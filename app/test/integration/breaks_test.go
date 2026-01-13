package integration_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"surf-share/app/internal/testutil"
	"testing"

	"surf-share/app/internal/adapters"
	"surf-share/app/internal/handlers"
	"surf-share/app/internal/models"
)

func TestGetBreaks(t *testing.T) {
	connStr, err := testutil.GetDbConnectionString()
	if err != nil {
		t.Fatalf("Failed to get DB connection string: %v", err)
	}

	ctx := context.Background()
	dbAdapter := &adapters.DatabaseAdapter{}
	if err := dbAdapter.Connect(ctx, connStr); err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbAdapter.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /breaks", handlers.NewBreaksHandler(dbAdapter).HandleBreaks)

	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + "/breaks")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Expected Content-Type application/json, got %s", contentType)
	}

	var response struct {
		Count  int                     `json:"count"`
		Breaks []models.BreaksResponse `json:"breaks"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Count != len(response.Breaks) || response.Count == 0 {
		t.Fatalf("Expected count %d to match breaks length %d and not be 0", response.Count, len(response.Breaks))
	}

	if response.Breaks[0].Slug != "la-arena" {
		t.Fatalf("Expected first break slug to be 'la-arena', got '%s'", response.Breaks[0].Slug)
	}

	t.Logf("✓ Successfully fetched %d breaks", response.Count)
}

func TestGetBreakBySlug(t *testing.T) {
	connStr, err := testutil.GetDbConnectionString()
	if err != nil {
		t.Fatalf("Failed to get DB connection string: %v", err)
	}

	ctx := context.Background()
	dbAdapter := &adapters.DatabaseAdapter{}
	if err := dbAdapter.Connect(ctx, connStr); err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbAdapter.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /breaks/{slug}", handlers.NewBreaksHandler(dbAdapter).HandleBreakBySlug)

	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := http.Get(server.URL + "/breaks/la-arena")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Fatalf("Expected Content-Type application/json, got %s", contentType)
	}

	var breakResponse models.BreakResponse
	if err := json.NewDecoder(resp.Body).Decode(&breakResponse); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if breakResponse.Slug != "la-arena" {
		t.Fatalf("Expected break slug to be 'la-arena', got '%s'", breakResponse.Slug)
	}

	if breakResponse.Name != "La Arena" {
		t.Fatalf("Expected break name to be 'La Arena', got '%s'", breakResponse.Name)
	}

	if breakResponse.Country != "ESP" {
		t.Fatalf("Expected break country to be 'ESP', got '%s'", breakResponse.Country)
	}

	t.Logf("✓ Successfully fetched break by slug: %s", breakResponse.Name)
}

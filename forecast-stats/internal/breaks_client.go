package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type BreakCoordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type BreakResponse struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Coordinates BreakCoordinates `json:"coordinates"`
	Country     string           `json:"country"`
}

type rawBreakResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Coordinates string `json:"coordinates"`
	Country     string `json:"country"`
}

type BreaksClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewBreaksClient() *BreaksClient {
	baseURL := os.Getenv("APP_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://app:8080"
	}

	return &BreaksClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *BreaksClient) GetBreakBySlug(ctx context.Context, slug string) (*BreakResponse, error) {
	url := fmt.Sprintf("%s/breaks/%s", c.baseURL, slug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch break: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	var rawResp rawBreakResponse
	if err := json.NewDecoder(resp.Body).Decode(&rawResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	coords, err := parseCoordinates(rawResp.Coordinates)
	if err != nil {
		return nil, fmt.Errorf("failed to parse coordinates: %w", err)
	}

	return &BreakResponse{
		ID:          rawResp.ID,
		Name:        rawResp.Name,
		Slug:        rawResp.Slug,
		Coordinates: coords,
		Country:     rawResp.Country,
	}, nil
}

func parseCoordinates(coordStr string) (BreakCoordinates, error) {
	coordStr = strings.Trim(coordStr, "()")
	parts := strings.Split(coordStr, ",")
	if len(parts) != 2 {
		return BreakCoordinates{}, fmt.Errorf("invalid coordinate format: %s", coordStr)
	}

	y, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return BreakCoordinates{}, err
	}

	x, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return BreakCoordinates{}, err
	}

	return BreakCoordinates{X: x, Y: y}, nil
}

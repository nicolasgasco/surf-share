package forecast

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenMeteoClient struct {
	baseURL string
	client  *http.Client
}

func NewOpenMeteoClient() *OpenMeteoClient {
	return &OpenMeteoClient{
		baseURL: "https://marine-api.open-meteo.com/v1/marine",
		client:  &http.Client{},
	}
}

type MarineForecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time                  []string  `json:"time"`
		WaveHeight            []float64 `json:"wave_height"`
		WavePeriod            []float64 `json:"wave_period"`
		WaveDirection         []int     `json:"wave_direction"`
		SeaSurfaceTemperature []float64 `json:"sea_surface_temperature"`
		SeaLevelHeightMsl     []float64 `json:"sea_level_height_msl"`
	} `json:"hourly"`
}

func (c *OpenMeteoClient) GetMarineForecast(ctx context.Context, latitude, longitude float64) (*MarineForecast, error) {
	includedParameters := "wave_height,wave_period,wave_direction,sea_surface_temperature,sea_level_height_msl"
	url := fmt.Sprintf("%s?latitude=%.4f&longitude=%.4f&hourly=%s&timezone=Europe%%2FBerlin&forecast_days=3",
		c.baseURL, latitude, longitude, includedParameters)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	var forecast MarineForecast
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, err
	}

	return &forecast, nil
}

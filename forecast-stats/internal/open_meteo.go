package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenMeteoClient struct {
	marineBaseURL string
	client        *http.Client
}

func NewOpenMeteoClient() *OpenMeteoClient {
	return &OpenMeteoClient{
		marineBaseURL: "https://marine-api.open-meteo.com/v1/marine",
		client:        &http.Client{},
	}
}

type MarineForecast struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	GenerationTime float64 `json:"generationtime_ms"`
	Timezone       string  `json:"timezone"`
	Elevation      float64 `json:"elevation"`
	HourlyUnits    struct {
		Time                  string `json:"time"`
		WaveHeight            string `json:"wave_height"`
		WavePeriod            string `json:"wave_period"`
		WaveDirection         string `json:"wave_direction"`
		SeaSurfaceTemperature string `json:"sea_surface_temperature"`
		SeaLevelHeightMsl     string `json:"sea_level_height_msl"`
	} `json:"hourly_units"`
	DailyUnits struct {
		Time                  string `json:"time"`
		WaveHeightMax         string `json:"wave_height_max"`
		WaveDirectionDominant string `json:"wave_direction_dominant"`
		WavePeriodMax         string `json:"wave_period_max"`
	} `json:"daily_units"`
	Hourly struct {
		Time                  []string  `json:"time"`
		WaveHeight            []float64 `json:"wave_height"`
		WavePeriod            []float64 `json:"wave_period"`
		WaveDirection         []int     `json:"wave_direction"`
		SeaSurfaceTemperature []float64 `json:"sea_surface_temperature"`
		SeaLevelHeightMsl     []float64 `json:"sea_level_height_msl"`
	} `json:"hourly"`
	Daily struct {
		Time                  []string  `json:"time"`
		WaveHeightMax         []float64 `json:"wave_height_max"`
		WaveDirectionDominant []int     `json:"wave_direction_dominant"`
		WavePeriodMax         []float64 `json:"wave_period_max"`
	} `json:"daily"`
}

func (c *OpenMeteoClient) GetMarineForecast(ctx context.Context, latitude, longitude float64) (*MarineForecast, error) {
	includedParameters := "wave_height,wave_direction,wave_period,sea_level_height_msl,sea_surface_temperature"
	url := fmt.Sprintf("%s?latitude=%.4f&longitude=%.4f&daily=wave_height_max,wave_direction_dominant,wave_period_max&hourly=%s&timezone=Europe%%2FBerlin&past_days=92&forecast_days=1",
		c.marineBaseURL, latitude, longitude, includedParameters)

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

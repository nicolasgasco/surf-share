package internal

import (
	"context"
)

type Stats struct {
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

type StatsService interface {
	GetStats(ctx context.Context, slug string) (*Stats, error)
}

func NewStatsService(openMeteoClient *OpenMeteoClient) StatsService {
	return &statsService{
		openMeteoClient: openMeteoClient,
	}
}

type statsService struct {
	openMeteoClient *OpenMeteoClient
}

func (s *statsService) GetStats(ctx context.Context, slug string) (*Stats, error) {
	// TODO: Hardcoded coordinates for now - will be improved later
	latitude := 43.458336
	longitude := -3.1249847

	forecast, err := s.openMeteoClient.GetMarineForecast(ctx, latitude, longitude)
	if err != nil {
		return nil, err
	}

	stats := &Stats{
		Latitude:       forecast.Latitude,
		Longitude:      forecast.Longitude,
		GenerationTime: forecast.GenerationTime,
		Timezone:       forecast.Timezone,
		Elevation:      forecast.Elevation,
		HourlyUnits:    forecast.HourlyUnits,
		DailyUnits:     forecast.DailyUnits,
		Hourly:         forecast.Hourly,
		Daily:          forecast.Daily,
	}

	return stats, nil
}

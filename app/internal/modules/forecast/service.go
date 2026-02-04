package forecast

import (
	"context"
	"fmt"
)

type FullForecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time                  []string  `json:"time"`
		WaveHeight            []float64 `json:"waveHeight"`
		WavePeriod            []float64 `json:"wavePeriod"`
		WaveDirection         []int     `json:"waveDirection"`
		SeaSurfaceTemperature []float64 `json:"seaSurfaceTemperature"`
		SeaLevelHeightMsl     []float64 `json:"seaLevelHeightMsl"`
		Temperature2m         []float64 `json:"temperature2m"`
		WindSpeed10m          []float64 `json:"windSpeed10m"`
		WindDirection10m      []int     `json:"windDirection10m"`
	} `json:"hourly"`
}

type StatsService interface {
	GetForecast(ctx context.Context, slug string) (*FullForecast, error)
}

func NewStatsService(repo StatsRepository, openMeteoClient *OpenMeteoClient) StatsService {
	return &statsService{
		repo:            repo,
		openMeteoClient: openMeteoClient,
	}
}

type statsService struct {
	repo            StatsRepository
	openMeteoClient *OpenMeteoClient
}

func (s *statsService) GetForecast(ctx context.Context, slug string) (*FullForecast, error) {
	coords, err := s.repo.GetBreakCoordinatesBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	marineForecast, err := s.openMeteoClient.GetMarineForecast(ctx, coords.Latitude, coords.Longitude)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch marine forecast: %w", err)
	}

	weatherForecast, err := s.openMeteoClient.GetWeatherForecast(ctx, coords.Latitude, coords.Longitude)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather forecast: %w", err)
	}

	fullForecast := &FullForecast{
		Latitude:  marineForecast.Latitude,
		Longitude: marineForecast.Longitude,
	}

	fullForecast.Hourly.Time = marineForecast.Hourly.Time
	fullForecast.Hourly.WaveHeight = marineForecast.Hourly.WaveHeight
	fullForecast.Hourly.WavePeriod = marineForecast.Hourly.WavePeriod
	fullForecast.Hourly.WaveDirection = marineForecast.Hourly.WaveDirection
	fullForecast.Hourly.SeaSurfaceTemperature = marineForecast.Hourly.SeaSurfaceTemperature
	fullForecast.Hourly.SeaLevelHeightMsl = marineForecast.Hourly.SeaLevelHeightMsl
	fullForecast.Hourly.Temperature2m = weatherForecast.Hourly.Temperature2m
	fullForecast.Hourly.WindSpeed10m = weatherForecast.Hourly.WindSpeed10m
	fullForecast.Hourly.WindDirection10m = weatherForecast.Hourly.WindDirection10m

	return fullForecast, nil
}

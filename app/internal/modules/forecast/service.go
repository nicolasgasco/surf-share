package forecast

import (
	"context"
	"fmt"
)

type StatsService interface {
	GetForecast(ctx context.Context, slug string) (*MarineForecast, error)
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

func (s *statsService) GetForecast(ctx context.Context, slug string) (*MarineForecast, error) {
	coords, err := s.repo.GetBreakCoordinatesBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	forecast, err := s.openMeteoClient.GetMarineForecast(ctx, coords.Latitude, coords.Longitude)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch marine forecast: %w", err)
	}

	return forecast, nil
}

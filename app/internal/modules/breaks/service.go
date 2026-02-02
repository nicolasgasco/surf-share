package breaks

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type breaksResponse struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
}

type breakResponse struct {
	ID          pgtype.UUID  `json:"id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Description pgtype.Text  `json:"description"`
	Coordinates pgtype.Point `json:"coordinates"`
	Country     string       `json:"country"`
	Region      pgtype.Text  `json:"region"`
	City        pgtype.Text  `json:"city"`
	VideoUrl    pgtype.Text  `json:"video_url"`
	ImageUrls   []string     `json:"image_urls"`
}

type BreaksService interface {
	GetBreaks(ctx context.Context) ([]breaksResponse, error)
	GetBreakBySlug(ctx context.Context, slug string) (*breakResponse, error)
}

func NewBreaksService(repo BreaksRepository) BreaksService {
	return &breaksService{
		repo: repo,
	}
}

type breaksService struct {
	repo BreaksRepository
}

type serviceBreaksResponse struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
}

type serviceBreakResponse struct {
	ID          pgtype.UUID  `json:"id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Description pgtype.Text  `json:"description"`
	Coordinates pgtype.Point `json:"coordinates"`
	Country     string       `json:"country"`
	Region      pgtype.Text  `json:"region"`
	City        pgtype.Text  `db:"city"`
	VideoUrl    pgtype.Text  `json:"video_url"`
	ImageUrls   []string     `json:"image_urls"`
}

func (s *breaksService) GetBreaks(ctx context.Context) ([]breaksResponse, error) {
	dbBreaks, err := s.repo.GetBreaks(ctx)
	if err != nil {
		return nil, err
	}

	breaks := make([]breaksResponse, len(dbBreaks))
	for i, db := range dbBreaks {
		breaks[i] = breaksResponse{
			ID:   db.ID,
			Name: db.Name,
			Slug: db.Slug,
		}
	}

	return breaks, nil
}

func (s *breaksService) GetBreakBySlug(ctx context.Context, slug string) (*breakResponse, error) {
	dbBreak, err := s.repo.GetBreakBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &breakResponse{
		ID:          dbBreak.ID,
		Name:        dbBreak.Name,
		Slug:        dbBreak.Slug,
		Description: dbBreak.Description,
		Coordinates: dbBreak.Coordinates,
		Country:     dbBreak.Country,
		Region:      dbBreak.Region,
		City:        dbBreak.City,
		VideoUrl:    dbBreak.VideoUrl,
		ImageUrls:   dbBreak.ImageUrls,
	}, nil
}

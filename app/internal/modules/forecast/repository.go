package forecast

import (
	"context"

	"surf-share/app/internal/adapters"

	"github.com/jackc/pgx/v5/pgtype"
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type StatsRepository interface {
	GetBreakCoordinatesBySlug(ctx context.Context, slug string) (*Coordinates, error)
}

type Repository struct {
	db *adapters.DatabaseAdapter
}

func NewRepository(db *adapters.DatabaseAdapter) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBreakCoordinatesBySlug(ctx context.Context, slug string) (*Coordinates, error) {
	var point pgtype.Point
	err := r.db.FindOne(
		ctx,
		&point,
		"SELECT coordinates FROM app.breaks WHERE slug = $1",
		slug,
	)
	if err != nil {
		return nil, err
	}

	return &Coordinates{
		Latitude:  point.P.X,
		Longitude: point.P.Y,
	}, nil
}

package breaks

import (
	"context"

	"surf-share/app/internal/adapters"

	"github.com/jackc/pgx/v5/pgtype"
)

type repoBreaksRow struct {
	ID   pgtype.UUID `db:"id"`
	Name string      `db:"name"`
	Slug string      `db:"slug"`
}

type repoBreakRow struct {
	ID          pgtype.UUID  `db:"id"`
	Name        string       `db:"name"`
	Slug        string       `db:"slug"`
	Description pgtype.Text  `db:"description"`
	Coordinates pgtype.Point `db:"coordinates"`
	Country     string       `db:"country"`
	Region      pgtype.Text  `db:"region"`
	City        pgtype.Text  `db:"city"`
	VideoUrl    pgtype.Text  `db:"video_url"`
	ImageUrls   []string     `db:"image_urls"`
}

type BreaksRepository interface {
	GetBreaks(ctx context.Context) ([]repoBreaksRow, error)
	GetBreakBySlug(ctx context.Context, slug string) (*repoBreakRow, error)
}

type Repository struct {
	db *adapters.DatabaseAdapter
}

func NewRepository(db *adapters.DatabaseAdapter) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBreaks(ctx context.Context) ([]repoBreaksRow, error) {
	breaks := make([]repoBreaksRow, 0)

	if err := r.db.FindMany(ctx, &breaks, "SELECT id, name, slug FROM app.breaks ORDER BY name ASC"); err != nil {
		return nil, err
	}

	return breaks, nil
}

func (r *Repository) GetBreakBySlug(ctx context.Context, slug string) (*repoBreakRow, error) {
	var brk repoBreakRow
	err := r.db.FindOne(
		ctx,
		&brk,
		`SELECT b.id, b.name, b.slug, b.description, b.coordinates, b.country, b.region, b.city, m.video_url, m.image_urls
		 FROM app.breaks b
		 LEFT JOIN app.breaks_media m ON b.slug = m.break_slug
		 WHERE b.slug = $1`,
		slug,
	)

	if err != nil {
		return nil, err
	}

	return &brk, nil
}

package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type BreakSummary struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
}

type Break struct {
	BreakSummary
	Description pgtype.Text  `json:"description"`
	Coordinates pgtype.Point `json:"coordinates"`
	Country     string       `json:"country"`
	Region      pgtype.Text  `json:"region"`
	City        pgtype.Text  `json:"city"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	VideoUrl    pgtype.Text  `json:"video_url"`
	ImageUrls   []string     `json:"image_urls"`
}

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BreaksResponse struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
}

type BreakResponse struct {
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

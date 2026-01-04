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
	Description pgtype.Text  `json:"description,omitempty"`
	Coordinates pgtype.Point `json:"coordinates"`
	Country     string       `json:"country"`
	Region      pgtype.Text  `json:"region,omitempty"`
	City        pgtype.Text  `json:"city,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

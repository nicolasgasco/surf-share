package auth

import (
	"context"

	"surf-share/app/internal/adapters"
)

type UserWithPassword struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type DbUser struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

type UserAuthRepository interface {
	FindUserWithPasswordByEmail(ctx context.Context, email string) (*UserWithPassword, error)
	CreateUser(ctx context.Context, username, email, hashedPassword string) (*DbUser, error)
}

type Repository struct {
	db *adapters.DatabaseAdapter
}

func NewRepository(db *adapters.DatabaseAdapter) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUserWithPasswordByEmail(ctx context.Context, email string) (*UserWithPassword, error) {
	var userWithPassword UserWithPassword

	if err := r.db.FindOne(ctx, &userWithPassword,
		"SELECT id, username, email, password FROM app.users WHERE email = $1",
		email); err != nil {
		return nil, err
	}

	return &userWithPassword, nil
}

func (r *Repository) CreateUser(ctx context.Context, username, email, hashedPassword string) (*DbUser, error) {
	var user DbUser

	if err := r.db.CreateOne(ctx, &user,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword); err != nil {
		return nil, err
	}

	return &user, nil
}

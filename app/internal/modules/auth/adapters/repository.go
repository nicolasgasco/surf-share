package adapters

import (
	"context"

	"surf-share/app/internal/adapters"
	"surf-share/app/internal/modules/auth"
)

type UserRepository struct {
	db *adapters.DatabaseAdapter
}

func NewUserRepository(db *adapters.DatabaseAdapter) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserCredentialsByEmail(ctx context.Context, email string) (*auth.UserCredentials, error) {
	var userCredentials auth.UserCredentials
	if err := r.db.FindOne(ctx, &userCredentials,
		"SELECT id, username, email, password FROM app.users WHERE email = $1",
		email); err != nil {
		return nil, err
	}
	return &userCredentials, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*auth.User, error) {
	var user auth.User
	if err := r.db.FindOne(ctx, &user,
		"SELECT id, username, email FROM app.users WHERE id = $1",
		id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, username, email, hashedPassword string) (*auth.User, error) {
	var user auth.User
	if err := r.db.CreateOne(ctx, &user,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword); err != nil {
		return nil, err
	}
	return &user, nil
}

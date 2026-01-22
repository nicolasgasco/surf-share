package auth

import (
	"context"

	"surf-share/app/internal/modules/auth/adapters"
)

type UserAuthRepository interface {
	FindUserCredentialsByEmail(ctx context.Context, email string) (*adapters.UserCredentials, error)
	FindUserByID(ctx context.Context, id string) (*adapters.User, error)
	CreateUser(ctx context.Context, username, email, hashedPassword string) (*adapters.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, password string) error
}

type TokenGenerator interface {
	Generate(user *adapters.User) (string, error)
}

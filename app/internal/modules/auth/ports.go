package auth

import "context"

type UserAuthRepository interface {
	FindUserCredentialsByEmail(ctx context.Context, email string) (*UserCredentials, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, username, email, hashedPassword string) (*User, error)
}

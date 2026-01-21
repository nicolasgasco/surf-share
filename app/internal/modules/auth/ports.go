package auth

import "context"

type UserAuthRepository interface {
	FindUserCredentialsByEmail(ctx context.Context, email string) (*UserCredentials, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, username, email, hashedPassword string) (*User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, password string) error
}

type TokenGenerator interface {
	Generate(user *User) (string, error)
}

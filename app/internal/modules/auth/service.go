package auth

import (
	"context"
	"strings"

	"surf-share/app/internal/modules/auth/adapters"
)

type AuthService struct {
	repo           UserAuthRepository
	passwordHasher PasswordHasher
	tokenGenerator TokenGenerator
}

func NewAuthService(repo UserAuthRepository, hasher PasswordHasher, generator TokenGenerator) *AuthService {
	return &AuthService{
		repo:           repo,
		passwordHasher: hasher,
		tokenGenerator: generator,
	}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (*adapters.User, string, error) {
	email = strings.ToLower(email)

	hashedPassword, err := s.passwordHasher.Hash(password)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, "", err
	}

	token, err := s.tokenGenerator.Generate(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*adapters.User, string, error) {
	email = strings.ToLower(email)

	userCredentials, err := s.repo.FindUserCredentialsByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}

	if err := s.passwordHasher.Verify(userCredentials.Password, password); err != nil {
		return nil, "", err
	}

	user, err := s.repo.FindUserByID(ctx, userCredentials.ID)
	if err != nil {
		return nil, "", err
	}

	token, err := s.tokenGenerator.Generate(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

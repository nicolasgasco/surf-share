package auth

import (
	"context"
	"strings"
)

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*AuthResult, error)
	Login(ctx context.Context, email, password string) (*AuthResult, error)
}

func NewAuthService(repo UserAuthRepository, hasher PasswordHasher, tokenGenerator TokenGenerator) AuthService {
	return &authService{
		repo:           repo,
		passwordHasher: hasher,
		tokenGenerator: tokenGenerator,
	}
}

type authService struct {
	repo           UserAuthRepository
	passwordHasher PasswordHasher
	tokenGenerator TokenGenerator
}

func (s *authService) Register(ctx context.Context, username, email, password string) (*AuthResult, error) {
	email = strings.ToLower(email)

	hashedPassword, err := s.passwordHasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenGenerator.Generate(user)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (*AuthResult, error) {
	email = strings.ToLower(email)

	userWithPassword, err := s.repo.FindUserWithPasswordByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := s.passwordHasher.Verify(userWithPassword.Password, password); err != nil {
		return nil, err
	}

	user := &User{
		ID:       userWithPassword.ID,
		Username: userWithPassword.Username,
		Email:    userWithPassword.Email,
	}

	token, err := s.tokenGenerator.Generate(user)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

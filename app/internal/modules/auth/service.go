package auth

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo UserAuthRepository
}

func NewAuthService(repo UserAuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (*User, string, error) {
	email = strings.ToLower(email)

	hashedPassword, err := s.encryptPassword(password)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.CreateUser(ctx, username, email, hashedPassword)
	if err != nil {
		return nil, "", err
	}

	token, err := s.createJwtToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*User, string, error) {
	email = strings.ToLower(email)

	userCredentials, err := s.repo.FindUserCredentialsByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}

	if err := s.verifyPassword(userCredentials.Password, password); err != nil {
		return nil, "", err
	}

	user, err := s.repo.FindUserByID(ctx, userCredentials.ID)
	if err != nil {
		return nil, "", err
	}

	token, err := s.createJwtToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *AuthService) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *AuthService) createJwtToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"iat":      time.Now(),
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

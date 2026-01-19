package auth

import (
	"context"
	"os"
	"strings"
	"time"

	"surf-share/app/internal/adapters"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewAuthService(dbAdapter *adapters.DatabaseAdapter) *AuthService {
	return &AuthService{dbAdapter: dbAdapter}
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) (*User, string, error) {
	email = strings.ToLower(email)

	hashedPassword, err := s.encryptPassword(password)
	if err != nil {
		return nil, "", err
	}

	var user User
	err = s.dbAdapter.CreateOne(ctx, &user,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword)
	if err != nil {
		return nil, "", err
	}

	token, err := s.createJwtToken(&user)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
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

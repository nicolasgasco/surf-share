package adapters

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

type UserCredentials struct {
	ID       string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type JWTGenerator struct {
	secret []byte
}

func NewJWTGenerator(secret []byte) *JWTGenerator {
	return &JWTGenerator{secret: secret}
}

func (j *JWTGenerator) Generate(user *User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"surf-share/app/internal/adapters"
)

type AuthHandler struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewAuthHandler(dbAdapter *adapters.DatabaseAdapter) *AuthHandler {
	return &AuthHandler{dbAdapter: dbAdapter}
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	lowercaseEmail := strings.ToLower(email)

	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	hashedPassword, err := encryptPassword(password)
	if err != nil {
		http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
		return
	}

	var newUser User
	ctx := context.Background()
	err = h.dbAdapter.CreateOne(ctx, &newUser,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, lowercaseEmail, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusBadRequest)
		return
	}

	token, err := createJwtToken(newUser.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  newUser,
		"token": token,
	}); err != nil {
		http.Error(w, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}

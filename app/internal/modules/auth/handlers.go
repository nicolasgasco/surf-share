package auth

import (
	"context"
	"net/http"
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

	ctx := context.Background()
	err = h.dbAdapter.Exec(ctx,
		"INSERT INTO app.users (username, email, password) VALUES ($1, $2, $3)",
		username, email, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("User registered")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"surf-share/app/internal/adapters"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(dbAdapter *adapters.DatabaseAdapter) *AuthHandler {
	return &AuthHandler{service: NewAuthService(dbAdapter)}
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

	ctx := context.Background()
	user, token, err := h.service.Register(ctx, username, email, password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()
	user, token, err := h.service.Login(ctx, email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

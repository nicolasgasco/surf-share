package auth

import (
	"context"
	"encoding/json"
	"net/http"
)

type HTTPHandler struct {
	authService AuthService
}

func NewHTTPHandler(authService AuthService) *HTTPHandler {
	return &HTTPHandler{
		authService: authService,
	}
}

func (h *HTTPHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
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
	result, err := h.authService.Register(ctx, username, email, password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  result.User,
		"token": result.Token,
	})
}

func (h *HTTPHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
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
	result, err := h.authService.Login(ctx, email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  result.User,
		"token": result.Token,
	})
}

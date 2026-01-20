package auth

import "net/http"

type AuthModule struct {
	service *AuthService
}

func NewAuthModule(repo UserAuthRepository) *AuthModule {
	service := NewAuthService(repo)
	return &AuthModule{service: service}
}

func (m *AuthModule) Register(mux *http.ServeMux) {
	authHandler := NewAuthHandler(m.service)

	mux.HandleFunc("POST /auth/login", authHandler.HandleLogin)
	mux.HandleFunc("POST /auth/register", authHandler.HandleRegister)
}

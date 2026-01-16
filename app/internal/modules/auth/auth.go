package auth

import (
	"net/http"
	"surf-share/app/internal/adapters"
)

type AuthModule struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewAuthModule(dbAdapter *adapters.DatabaseAdapter) *AuthModule {
	return &AuthModule{dbAdapter: dbAdapter}
}

func (m *AuthModule) Register(mux *http.ServeMux) {
	authHandler := NewAuthHandler(m.dbAdapter)
	mux.HandleFunc("POST /auth/register", authHandler.HandleRegister)
}

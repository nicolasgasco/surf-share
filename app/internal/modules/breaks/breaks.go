package breaks

import (
	"net/http"
	"surf-share/app/internal/adapters"
)

type BreaksModule struct {
	dbAdapter *adapters.DatabaseAdapter
}

func NewBreaksModule(dbAdapter *adapters.DatabaseAdapter) *BreaksModule {
	return &BreaksModule{dbAdapter: dbAdapter}
}

func (m *BreaksModule) Register(mux *http.ServeMux) {
	breaksHandler := NewBreaksHandler(m.dbAdapter)
	mux.HandleFunc("GET /breaks", breaksHandler.HandleBreaks)
	mux.HandleFunc("GET /breaks/{slug}", breaksHandler.HandleBreakBySlug)
}

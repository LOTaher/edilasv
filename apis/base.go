package apis

import (
	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
)

// InitApi creates a configured chi router with necessary routes and middlewares
func InitAPI(db *core.Store) (*chi.Mux, error) {
	r := chi.NewRouter()

	// Middlewares
    // ...

	// Routes
	r.Mount("/api", CrudAPIRoutes(db))

	return r, nil
}

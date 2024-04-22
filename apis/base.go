package apis

import (
	"github.com/go-chi/chi/v5"
)

// InitApi creates a configured chi router with necessary routes and middlewares
func InitApi() (*chi.Mux, error) {
	r := chi.NewRouter()

	// Middlewares
	// r.Use()

	// Routes
	r.Mount("/api", CrudApiRoutes())

	return r, nil
}

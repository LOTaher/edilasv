package apis

import (
	"encoding/json"
	"net/http"

	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// the endpoint response function
func SendJSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Status:  status,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func InitAPI(db *core.Store, config ServeConfig) (*chi.Mux, error) {
	r := chi.NewRouter()

	// global middleware
	// cors configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

    // database store middleware
	r.Use(DatabaseMiddleware(db))

	// api key validation
	r.Use(KeyMiddleware(config.Key))

	// routes
	// the record CRUD api
	r.Mount("/api", CrudRoutes())

	return r, nil
}

package apis

import (
    "encoding/json"
    "net/http"

	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
)

type Response struct {
    Status int `json:"status"`
    Message string `json:"message"`
}

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

func InitAPI(db *core.Store) (*chi.Mux, error) {
	r := chi.NewRouter()

	// Middlewares
    // ...

	// Routes
	r.Mount("/api", CrudRoutes(db))

	return r, nil
}

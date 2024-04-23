package apis

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CrudApiRoutes returns the routes for the CRUD API
func CrudApiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{key}", GetKVPair)
	r.Post("/{key}", CreateKVPair)
	r.Put("/{key}", UpdateKVPair)
	r.Delete("/{key}", DeleteKVPair)

	return r
}

// GetKVPair returns the value for a given key
func GetKVPair(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// CreateKVPair creates a new key-value pair
func CreateKVPair(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// UpdateKVPair updates the value for a given key
func UpdateKVPair(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// DeleteKVPair deletes a key-value pair
func DeleteKVPair(w http.ResponseWriter, r *http.Request) {
	// TODO
}

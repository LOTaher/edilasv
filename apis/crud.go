package apis

import (
	"encoding/json"
	"net/http"

	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
)

// CrudAPIRoutes returns the routes for the CRUD API
func CrudAPIRoutes(db *core.Store) chi.Router {
	r := chi.NewRouter()

	r.Use(DatabaseMiddleware(db))

	r.Get("/get/{key}", GetKVPair)
	r.Post("/insert", InsertKVPair)
	r.Put("/update", UpdateKVPair)
	r.Delete("/delete/{key}", DeleteKVPair)

	return r
}

// GetKVPair returns the value for a given key
func GetKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreContextKey{}).(*core.Store)

	var item core.Item
	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value, ok := store.Get(key)
	if ok {
		item = core.Item{Key: key, Value: value}
		json.NewEncoder(w).Encode(item)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// CreateKVPair creates a new key-value pair
func InsertKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreContextKey{}).(*core.Store)

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)

	if item.Key == "" || item.Value == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store.Insert(item.Key, item.Value)
	w.WriteHeader(http.StatusOK)
}

// UpdateKVPair updates the value for a given key
func UpdateKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreContextKey{}).(*core.Store)

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)

	if item.Key == "" || item.Value == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		store.Update(item.Key, item.Value)
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteKVPair deletes a key-value pair
func DeleteKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreContextKey{}).(*core.Store)

	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	store.Delete(key)
	w.WriteHeader(http.StatusOK)
}

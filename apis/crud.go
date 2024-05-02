package apis

import (
	"encoding/json"
	"net/http"

	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
)

func CrudRoutes(db *core.Store) chi.Router {
	r := chi.NewRouter()

	r.Use(DatabaseMiddleware(db))

	r.Get("/get/{key}", GetKVPair)
	r.Post("/insert", InsertKVPair)
	r.Put("/update", UpdateKVPair)
	r.Delete("/delete/{key}", DeleteKVPair)

	return r
}

func GetKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "Store not configured", nil)
		return
	}

	key := chi.URLParam(r, "key")
	if key == "" {
        SendJSONResponse(w, http.StatusBadRequest, "Key must be provided", nil)
		return
	}

	value, found := store.Get(key)
	if !found {
        SendJSONResponse(w, http.StatusNotFound, "Key not found", nil)
		return
	}

	item := core.Item{Key: key, Value: value}
    SendJSONResponse(w, http.StatusOK, "Successfully retrieved key-value pair", item)
}

func InsertKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "Store not configured", nil)
		return
	}

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)
	if item.Key == "" || item.Value == nil {
        SendJSONResponse(w, http.StatusBadRequest, "Key and value must be provided", nil)
		return
	}

	store.Insert(item.Key, item.Value)
    SendJSONResponse(w, http.StatusOK, "Successfully inserted key-value pair", nil)
}

func UpdateKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
		http.Error(w, "Store not configured", http.StatusInternalServerError)
		return
	}

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)

	if item.Key == "" || item.Value == nil {
        SendJSONResponse(w, http.StatusBadRequest, "Key and value must be provided", nil)
		return
	}

    if _, found := store.Get(item.Key); !found {
        SendJSONResponse(w, http.StatusNotFound, "Key not found", nil)
        return
    }

	store.Update(item.Key, item.Value)
    SendJSONResponse(w, http.StatusOK, "Successfully updated key-value pair", nil)
}

func DeleteKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "Store not configured", nil)
		return
	}

	key := chi.URLParam(r, "key")
	if key == "" {
        SendJSONResponse(w, http.StatusBadRequest, "Key must be provided", nil)
		return
	}
	store.Delete(key)
    SendJSONResponse(w, http.StatusOK, "Successfully deleted key-value pair", nil)
}

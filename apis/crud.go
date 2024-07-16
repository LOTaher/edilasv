package apis

import (
	"encoding/json"
	"net/http"

	"github.com/LOTaher/softbase/core"
	"github.com/go-chi/chi/v5"
)

func CrudRoutes(db *core.Store) chi.Router {
	r := chi.NewRouter()

    // middleware
	r.Use(DatabaseMiddleware(db))

    // routes
    r.Post("/create", CreateKVPair)
	r.Get("/read/{key}", ReadKVPair)
    r.Get("/read", GetAllKVPair)
	r.Put("/update", UpdateKVPair)
	r.Delete("/delete/{key}", DeleteKVPair)

	return r
}

func ReadKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "store not configured", nil)
		return
	}

	key := chi.URLParam(r, "key")
	if key == "" {
        SendJSONResponse(w, http.StatusBadRequest, "key must be provided", nil)
		return
	}

	value, found := store.Get(key)
	if !found {
        SendJSONResponse(w, http.StatusNotFound, "key not found", nil)
		return
	}

	item := core.Item{Key: key, Value: value}
    SendJSONResponse(w, http.StatusOK, "successfully read key-value pair", item)
}

func CreateKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "store not configured", nil)
		return
	}

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)
	if item.Key == "" || item.Value == nil {
        SendJSONResponse(w, http.StatusBadRequest, "key and value must be provided", nil)
		return
	}

    if store.Has(item.Key) {
        SendJSONResponse(w, http.StatusConflict, "key already exists", nil)
        return
    }

	store.Insert(item.Key, item.Value)
    SendJSONResponse(w, http.StatusOK, "successfully created key-value pair", nil)
}

func UpdateKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
		http.Error(w, "store not configured", http.StatusInternalServerError)
		return
	}

	var item core.Item
	json.NewDecoder(r.Body).Decode(&item)

	if item.Key == "" || item.Value == nil {
        SendJSONResponse(w, http.StatusBadRequest, "key and value must be provided", nil)
		return
	}

    if !store.Has(item.Key) {
        SendJSONResponse(w, http.StatusNotFound, "key not found", nil)
        return
    }

	store.Update(item.Key, item.Value)
    SendJSONResponse(w, http.StatusOK, "successfully updated key-value pair", nil)
}

func DeleteKVPair(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store, ok := ctx.Value(StoreContextKey).(*core.Store)
	if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "store not configured", nil)
		return
	}

	key := chi.URLParam(r, "key")
	if key == "" {
        SendJSONResponse(w, http.StatusBadRequest, "key must be provided", nil)
		return
	}

    if !store.Has(key) {
        SendJSONResponse(w, http.StatusNotFound, "key not found", nil)
        return
    }

	store.Delete(key)
    SendJSONResponse(w, http.StatusOK, "successfully deleted key-value pair", nil)
}

func GetAllKVPair(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    store, ok := ctx.Value(StoreContextKey).(*core.Store)
    if !ok {
        SendJSONResponse(w, http.StatusInternalServerError, "store not configured", nil)
        return
    }

    allItems := store.GetAll()
    items := make([]core.Item, 0, len(allItems));
    for _, item := range store.GetAll() {
        items = append(items, item)
    }

    SendJSONResponse(w, http.StatusOK, "successfully read all key-value pairs", items)
}

package http

import (
    "encoding/json"
    "net/http"

    "github.com/LOTaher/softbase/pkg/btree"
    "github.com/LOTaher/softbase/pkg/storage"
)

type Handler struct {
    Db storage.Database
}

func NewHandler(db storage.Database) *Handler {
    return &Handler{db}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    var item btree.Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    h.Db.AddItem(item.Key, item.Value)
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
    var item btree.Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    value, ok := h.Db.GetItem(item.Key)
    if !ok {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(value)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
    var item btree.Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    h.Db.DeleteItem(item.Key)
    w.WriteHeader(http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
    var item btree.Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    h.Db.UpdateItem(item.Key, item.Value)
    w.WriteHeader(http.StatusOK)
}

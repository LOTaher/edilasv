package http

import (
    "net/http"
    "log"
    "encoding/json"

    "github.com/LOTaher/softbase/pkg/btree"
)

type Handler struct {}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
    var item btree.Item

    // Need to find a way to pass the user's db to the handler
    // Finding the db should be in the start server function
    var db btree.Store

    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    db.Put(item.Key, item.Value)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
    var item btree.Item

    var db btree.Store
}


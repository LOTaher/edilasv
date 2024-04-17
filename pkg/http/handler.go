package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LOTaher/softbase/pkg/btree"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusOK)
	db.Insert(item.Key, item.Value)
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
	var item btree.Item

	var db btree.Store

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, false := db.Get(item.Key)
	if false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var item btree.Item

	var db btree.Store

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	db.Delete(item.Key)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var item btree.Item

	var db btree.Store

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	db.Update(item.Key, item.Value)
}

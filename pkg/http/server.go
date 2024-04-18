package http

import (
	"log"
	"net/http"

    "github.com/LOTaher/softbase/pkg/btree"
)

func StartServer(db *btree.Store) {
    handler := NewHandler(db)

	mux := http.NewServeMux()

    mux.HandleFunc("/api/create", handler.Create)
    mux.HandleFunc("/api/read", handler.Read)
    mux.HandleFunc("/api/delete", handler.Delete)
    mux.HandleFunc("/api/update", handler.Update)

	server := http.Server{
		Addr:    ":69420",
		Handler: mux,
    }

	log.Println("Serving SoftBase on http://localhost:69420")

	log.Fatal(server.ListenAndServe())
}

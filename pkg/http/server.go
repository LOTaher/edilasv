package http

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":69420",
		Handler: mux,
	}

	log.Println("Serving SoftBase on http://localhost:69420")

	log.Fatal(server.ListenAndServe())
}

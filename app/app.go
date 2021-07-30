package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Boot() {
	mux := mux.NewRouter()
	// routes
	mux.HandleFunc("/customers", Customers)
	// start server
	log.Fatal(http.ListenAndServe("localhost:7070", mux))
}

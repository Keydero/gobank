package app

import (
	"log"
	"net/http"
)

func Boot() {
	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/customers", Customers)
	// start server
	log.Fatal(http.ListenAndServe("localhost:7070", mux))
}

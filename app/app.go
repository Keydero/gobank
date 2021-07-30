package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Boot() {
	router := mux.NewRouter()
	// routes
	router.HandleFunc("/customers", Customers).Methods(http.MethodGet)
	// Get a Customer by ID
	router.HandleFunc("/customers/{id:[0-9]+}", GetCustomer).Methods(http.MethodGet)
	// Create a Customer
	router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	// start server
	log.Fatal(http.ListenAndServe("localhost:7070", router))
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	result := mux.Vars(r)
	fmt.Fprint(w, result["id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

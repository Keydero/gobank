package main

import (
	"net/http"
)

func Boot() {
	http.HandleFunc("/customers", Customers)
	http.ListenAndServe("localhost:7070", nil)
}

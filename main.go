package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)


type Customer struct {
	Name string `json:"fullName"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
}
func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/customers", customers)
	http.ListenAndServe("localhost:7070", nil)
}



func customers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{Name: "Simone", City:"Paris", Zipcode: "75010"},
		{Name: "Alicia", City:"Créteil", Zipcode: "94000"},
		{Name: "Ilies", City:"Alfortville", Zipcode: "94140"},
		{Name: "Romain", City:"Maison-Alfort", Zipcode: "94410"},
		{Name: "Celeste",City:"Ivry-Sur-Seine", Zipcode: "94205"},
		{Name: "Thomas", City:"Saint-Maur", Zipcode: "94100"},
		{Name: "Simone", City:"Orgrimmar", Zipcode: "44444"},
		{Name: "Estelle", City:"Lune d'Argent", Zipcode:"22222"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
package main
import (
	"fmt"
	"net/http" 
)
func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.ListenAndServe("localhost:7070", nil)
}


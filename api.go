package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola, probando codigo!")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			users := []map[string]string{
				{"name": "Jesus", "email": "jesus@example.com"},
				{"name": "Juan", "email": "juan@example.com"},
			}
			fmt.Fprint(w, users)
		}
	})
	http.ListenAndServe(":8080", nil)
}

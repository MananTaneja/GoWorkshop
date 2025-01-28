package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Unable to setup server on port 8080\n")
	}
	fmt.Printf("Server listening on port 8080")
}

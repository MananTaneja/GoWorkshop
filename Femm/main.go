package main

import (
	"fmt"
	"html/template"
	"net/http"

	"manantaneja.com/go/frontend-masters-museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.templ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to parse files"))
		return
	}

	err = html.Execute(w, data.GetAll()[0])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to parse html"))
		return
	}

}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Printf("Unable to setup server on port 8080\n")
	}
	fmt.Printf("Server listening on port 8080")
}

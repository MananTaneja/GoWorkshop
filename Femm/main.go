package main

import (
	"fmt"
	"html/template"
	"net/http"

	"manantaneja.com/go/frontend-masters-museum/api"
	"manantaneja.com/go/frontend-masters-museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("/Users/manan/Developer/GoWorkshop/Femm/templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to parse files"))
		return
	}

	err = html.Execute(w, data.GetAll())

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
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Printf("Unable to setup server on port 8080\n")
	}
	fmt.Printf("Server listening on port 8080")
}

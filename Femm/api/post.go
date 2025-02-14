package api

import (
	"encoding/json"
	"net/http"

	"manantaneja.com/go/frontend-masters-museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}
		data.AddExhibition(exhibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Exhibtion added"))
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}

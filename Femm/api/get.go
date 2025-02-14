package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"manantaneja.com/go/frontend-masters-museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		intId, err := strconv.Atoi(id[0])
		if err != nil || intId >= len(data.GetAll()) {
			http.Error(w, "Invalid id", http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(data.GetAll()[intId])
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}

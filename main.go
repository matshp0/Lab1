package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Time string `json:"time"`
	}{
		Time: time.Now().UTC().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)

	port := ":8795"
	http.ListenAndServe(port, nil)
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pingpongHandler)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func pingpongHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	}
	json.NewEncoder(w).Encode(&resp)
}

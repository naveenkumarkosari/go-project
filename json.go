package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}, message string) {
	w.Header().Set("Content-Type", "application.json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("error formating payload")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}

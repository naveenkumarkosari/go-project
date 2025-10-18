package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with error", msg)
	}
	type ErrorMsg struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, ErrorMsg{Error: msg})
}

func responseWithJSON(w http.ResponseWriter, code int, payload any) {
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

package main

import "net/http"

type body struct {
	Message string
	Status  bool
}

func handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, body{Message: "Health is Good", Status: true})
}

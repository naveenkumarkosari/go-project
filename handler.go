package main

import "net/http"

type body struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, body{Message: "Health is Good", Status: true})
}

func handleErrRequest(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "something went wrong")
}

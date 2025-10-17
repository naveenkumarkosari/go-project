package main

import "net/http"

type body struct {
	Message string
}

func handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, http.StatusOK, body{Message: "Health is Good"}, "ok")
}

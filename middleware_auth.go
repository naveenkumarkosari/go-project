package main

import (
	"net/http"

	"github.com/naveenkumarkosari/go-project.git/internal/auth"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 401, "unauthorized user")
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 401, "something went wrong fetching user")
			return
		}
		handler(w, r, user)
	}
}

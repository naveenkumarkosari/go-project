package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

func (apiCfg apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err, "error")
		responseWithError(w, 400, "Bad payload")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Createdat: time.Now(),
		Updatedat: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		fmt.Println("error creating User", err)
		responseWithError(w, 400, "Unable to create User")
	}
	newUser := databaseUserDyn(user)

	responseWithJSON(w, http.StatusOK, newUser)
}

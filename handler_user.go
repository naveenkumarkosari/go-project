package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/naveenkumarkosari/go-project.git/internal/auth"
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

func (apiCfg apiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 401, "unauthorized user")
		return
	}
	fmt.Println(apiKey, "===key")
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 401, "something went wrong fetching user")
		return
	}
	newUser := databaseUserDyn(user)
	responseWithJSON(w, 200, newUser)
}

func (apiCfg apiConfig) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if err != nil {
		responseWithError(w, 502, "Something went wrong")
		return
	}
	newUsers := []User{}
	for i := range len(users) {
		dynUser := databaseUserDyn(users[i])
		newUsers = append(newUsers, dynUser)
	}
	responseWithJSON(w, 200, newUsers)
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/naveenkumarkosari/go-project.git/internal/auth"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

func (apiCfg apiConfig) CreateFeed(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, "authenticated route")
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 401, "user not found")
	}
	type parameters struct {
		Text string `json:"text"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		fmt.Println(err, "parsing error")
		responseWithError(w, 400, "Bad payload")
		return
	}
	content := sql.NullString{
		String: params.Text,
		Valid:  params.Text != " ",
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Createdat: time.Now(),
		Updatedat: time.Now(),
		Content:   content,
		Createdby: user.ID,
	})
	if err != nil {
		responseWithError(w, 401, "something went wrong")
	}
	responseWithJSON(w, 200, feed)
}

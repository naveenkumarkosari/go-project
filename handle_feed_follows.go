package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

func (apiCfg apiConfig) CreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 403, "Bad Payload")
	}

	_, err = apiCfg.DB.GetFeedById(r.Context(), params.FeedID)
	if err != nil {
		responseWithError(w, 500, "something wrong with feed")
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		Createdat: time.Now(),
		Updatedat: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		fmt.Println(err)
		responseWithError(w, 500, "something went wrong")
	}
	responseWithJSON(w, 200, dynFeedFollows(feedFollow))
}

func (apiCfg apiConfig) GetUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollows, err := apiCfg.DB.GetUserFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 500, "something went wrong")
	}
	result := []uuid.UUID{}
	for i := range feedfollows {
		result = append(result, feedfollows[i].FeedID)
	}
	responseWithJSON(w, 200, result)
}

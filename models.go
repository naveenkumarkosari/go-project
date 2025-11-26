package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/naveenkumarkosari/go-project.git/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	APIKey    string    `json:"apiKey"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	Createdby uuid.UUID `json:"creatorId"`
	Text      string    `json:"content"`
}
type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"userId"`
	FeedID    uuid.UUID `json:"feedId"`
}

func databaseUserDyn(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func dynFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.Createdat,
		Updatedat: dbFeed.Updatedat,
		Text:      dbFeed.Content.String,
		Createdby: dbFeed.Createdby,
	}
}

func dynFeedFollows(dbFeedFollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.Createdat,
		Updatedat: dbFeedFollows.Updatedat,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.FeedID,
	}
}

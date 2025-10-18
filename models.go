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

func databaseUserDyn(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

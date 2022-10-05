package api

import (
	"log"
)

func getUserID() string {
	user, err := Client.CurrentUser(Ctx)
	if err != nil {
		log.Default().Fatalln("Error fetching current user:", err)
	}
	userID := user.User.ID
	return userID
}

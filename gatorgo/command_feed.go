package main

import (
	"context"
	"fmt"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerGetFeeds(s *config.State, cmd config.Cmd) error {
	// set context, get user
	ctxt := context.Background()
	feeds, errGetFeeds := s.Db_Queries.GetAllFeeds(ctxt)
	if errGetFeeds != nil {
		return errGetFeeds
	}
	for _, feed := range feeds {
		user, errGetUserById := s.Db_Queries.GetUserByID(ctxt, feed.UserID)
		if errGetUserById != nil {
			return errGetUserById
		}
		fmt.Printf("Feed name: %s \n Feed URL: %s \n Created by: %s \n", feed.Name, feed.Url, user.Name)
	}
	return nil
}

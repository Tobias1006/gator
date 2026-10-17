package main

import (
	"context"
	"fmt"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerFollowing(s *config.State, cmd config.Cmd) error {
	// set context, get user
	ctxt := context.Background()
	currentUser := s.Config.Current_user_name

	user, errGetUser := s.Db_Queries.GetUser(ctxt, currentUser)
	if errGetUser != nil {
		return errGetUser
	}
	feeds, errGetFeeds := s.Db_Queries.GetFeedFollowsForUser(ctxt, user.ID)
	if errGetFeeds != nil {
		return errGetFeeds
	}
	fmt.Printf("User: %s \n Feeds: \n", currentUser)
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\n", feed.FeedName)
	}
	return nil
}

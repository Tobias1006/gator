package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Tobias1006/gatorgo/internal/config"
	"github.com/Tobias1006/gatorgo/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *config.State, cmd config.Cmd) error {
	// checking if correct number of args was passed
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("incorrect number of arguments for the command found")
	}
	ctxt := context.Background()
	url := cmd.Arguments[0]

	feed, errGetFeed := s.Db_Queries.GetFeedByURL(ctxt, url)
	if errGetFeed != nil {
		return errGetFeed
	}
	user, errGetUser := s.Db_Queries.GetUser(ctxt, s.Config.Current_user_name)
	if errGetUser != nil {
		return errGetUser
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	followedFeed, errCreateFollow := s.Db_Queries.CreateFeedFollow(ctxt, feedFollowParams)
	if errCreateFollow != nil {
		return errCreateFollow
	}

	// message if successful
	fmt.Printf("Feedname: %s, User: %s", followedFeed.FeedName, followedFeed.UserName)
	return nil
}

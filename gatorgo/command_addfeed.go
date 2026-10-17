package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Tobias1006/gatorgo/internal/config"
	"github.com/Tobias1006/gatorgo/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *config.State, cmd config.Cmd) error {
	// checking if correct number of args was passed
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("incorrect number of arguments for the command found")
	}

	// set context, get user info
	ctxt := context.Background()
	currentUser, errGetUser := s.Db_Queries.GetUser(ctxt, s.Config.Current_user_name)
	if errGetUser != nil {
		return errGetUser
	}

	// creating context, assigning args
	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currentUser.ID,
	}

	// create new feed
	newFeed, errFeedCreation := s.Db_Queries.CreateFeed(ctxt, feedParams)
	if errFeedCreation != nil {
		return errFeedCreation
	}

	// create new feed follow
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    newFeed.ID,
		UserID:    currentUser.ID,
	}
	_, errFeedFollow := s.Db_Queries.CreateFeedFollow(ctxt, feedFollowParams)
	if errFeedFollow != nil {
		return errFeedFollow
	}

	// print feed details
	fmt.Printf("Feed %s was created. \n", newFeed.Name)
	fmt.Printf("F-Name: %v \n F-ID: %v \n F-CreatedAt: %v \n", newFeed.Name, newFeed.ID, newFeed.CreatedAt)
	fmt.Printf("F-UpdatedAt: %v \n F-URL: %v \n F-UserID: %v \n ", newFeed.UpdatedAt, newFeed.Url, newFeed.UserID)
	return nil
}

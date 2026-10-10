package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/Tobias1006/gatorgo/internal/config"
	"github.com/Tobias1006/gatorgo/internal/database"
)

func handlerRegister(s *config.State, cmd config.Cmd) error {
	// check input arguments
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("invalid number of arguments provided")
	}

	// define user to be added to the database
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	// query database to avoid adding an already existing user
	_, errGetUser := s.Db_Queries.GetUser(context.Background(), userParams.Name)
	if errGetUser == nil {
		os.Exit(1)
	} else if errors.Is(errGetUser, sql.ErrNoRows) { // check if error is expected sql error for missing rows
		// if user does not already exist: create User and print user data
		user, errCreateUser := s.Db_Queries.CreateUser(context.Background(), userParams)
		s.Config.SetUser(user.Name)
		fmt.Printf("User %s was created. \n", user.Name)
		fmt.Printf("U-Name: %v \n U-ID: %v \n U-CreatedAt: %v \n U-UpdatedAt: %v \n ", user.Name, user.ID, user.CreatedAt, user.UpdatedAt)
		return errCreateUser
	}
	return errGetUser
}

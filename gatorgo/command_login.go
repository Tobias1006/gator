package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerLogin(s *config.State, cmd config.Cmd) error {
	// checking if correct number of args was passed
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("incorrect number of arguments for the command found")
	}
	username := cmd.Arguments[0]
	_, errGetUser := s.Db_Queries.GetUser(context.Background(), username)
	if errGetUser != nil {
		os.Exit(1)
	}
	// setting user to provided user
	errSettingUser := s.Config.SetUser(username)
	if errSettingUser != nil {
		return errSettingUser
	}

	// message if successful
	fmt.Printf("User %s has been set successfully \n", username)
	return nil
}

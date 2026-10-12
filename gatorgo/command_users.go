package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func hanlderGetUsers(s *config.State, cmd config.Cmd) error {
	users, err := s.Db_Queries.GetAllUsers(context.Background())
	if err != nil {
		fmt.Print("Could not get all Users")
		os.Exit(1)
	}
	for i := range len(users) {
		if users[i].Name == s.Config.Current_user_name {
			fmt.Printf("* %s (current) \n", users[i].Name)
		} else {
			fmt.Printf("* %s \n", users[i].Name)
		}
	}
	return nil
}

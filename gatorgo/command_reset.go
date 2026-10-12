package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerReset(s *config.State, cmd config.Cmd) error {
	err := s.Db_Queries.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Print("Could not delete all users.\n")
		os.Exit(1)
	}
	fmt.Print("All users deleted. \n")
	return nil
}

package main

import (
	"fmt"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func handlerLogin(s *config.State, cmd config.Cmd) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("incorrect number of arguments for the command found")
	}
	errSettingUser := s.Config.SetUser(cmd.Arguments[0])
	if errSettingUser != nil {
		return errSettingUser
	}
	fmt.Printf("User %s has been set successfully \n", cmd.Arguments[0])
	return nil
}

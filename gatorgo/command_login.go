package main

import "fmt"

func handlerLogin(s *state, cmd cmd) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Incorrect number of arguments for the command found.")
	}
	errSettingUser := s.config.SetUser(cmd.arguments[0])
	if errSettingUser != nil {
		return errSettingUser
	}
	fmt.Printf("User %s has been set successfully", cmd.arguments[0])
	return nil
}

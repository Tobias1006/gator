package main

import (
	"fmt"

	"github.com/Tobias1006/gatorgo/internal/config"
)

type state struct {
	config *config.Config
}

type cmd struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, cmd) error
}

func (c *commands) run(s *state, cmd cmd) error {
	command, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command. \n")
	}
	errRunCommand := command(s, cmd)
	if errRunCommand != nil {
		return errRunCommand
	}
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd cmd) error) {
	_, ok := c.commandMap[name]
	if ok {
		fmt.Printf("Function %s is already registered \n", name)
		return
	}
	c.commandMap[name] = f
}

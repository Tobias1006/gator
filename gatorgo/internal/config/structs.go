package config

import (
	"fmt"
)

type State struct {
	Config *Config
}

type Cmd struct {
	Name      string
	Arguments []string
}

type Commands struct {
	CommandMap map[string]func(*State, Cmd) error
}

func (c *Commands) Run(s *State, cmd Cmd) error {
	command, ok := c.CommandMap[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command")
	}
	errRunCommand := command(s, cmd)
	if errRunCommand != nil {
		return errRunCommand
	}
	return nil
}

func (c *Commands) Register(name string, f func(s *State, cmd Cmd) error) {
	_, ok := c.CommandMap[name]
	if ok {
		fmt.Printf("Function %s is already registered \n", name)
		return
	}
	c.CommandMap[name] = f
}

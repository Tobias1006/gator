package main

import (
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func main() {
	var current_state config.State
	current_state.Config = config.Read()
	var cmds config.Commands
	cmds.CommandMap = make(map[string]func(*config.State, config.Cmd) error)
	cmds.Register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Print("Too few arguments. \n")
		os.Exit(1)
	}
	var current_command config.Cmd
	current_command.Name = args[1]
	current_command.Arguments = args[2:]
	errRun := cmds.Run(&current_state, current_command)
	if errRun != nil {
		fmt.Printf("%v \n", errRun)
		os.Exit(1)
	}
}

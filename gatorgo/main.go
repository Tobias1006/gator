package main

import (
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func main() {
	var current_state state
	current_state.config = config.Read()
	var cmds commands
	cmds.commandMap = make(map[string]func(*state, cmd) error)
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Print("Too few arguments. \n")
		os.Exit(1)
	}
	var current_command cmd
	current_command.name = args[1]
	current_command.arguments = args[2:]
	errRun := cmds.run(&current_state, current_command)
	if errRun != nil {
		fmt.Printf("%v \n", errRun)
		os.Exit(1)
	}
}

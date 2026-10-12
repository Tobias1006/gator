package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Tobias1006/gatorgo/internal/config"
	"github.com/Tobias1006/gatorgo/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// setting config state
	var current_state config.State
	current_state.Config = config.Read()

	// connecting DB
	db, errDB := sql.Open("postgres", current_state.Config.Db_url)
	if errDB != nil {
		fmt.Printf("%v \n", errDB)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	current_state.Db_Queries = dbQueries

	// getting and setting commands, reading user input
	var cmds config.Commands
	cmds.CommandMap = make(map[string]func(*config.State, config.Cmd) error)
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)
	cmds.Register("reset", handlerReset)
	cmds.Register("users", hanlderGetUsers)
	cmds.Register("agg", handlerAggregator)
	cmds.Register("addfeed", handlerAddFeed)
	cmds.Register("feeds", handlerGetFeeds)
	args := os.Args
	if len(args) < 2 {
		fmt.Print("Too few arguments. \n")
		os.Exit(1)
	}

	// extracting user-command from input
	var current_command config.Cmd
	current_command.Name = args[1]
	current_command.Arguments = args[2:]

	// running command given by user
	errRun := cmds.Run(&current_state, current_command)
	if errRun != nil {
		fmt.Printf("%v \n", errRun)
		os.Exit(1)
	}
}

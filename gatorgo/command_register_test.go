package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Tobias1006/gatorgo/internal/database"
	"github.com/Tobias1006/gatorgo/internal/testkit"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func TestHandlerRegister(t *testing.T) {
	fmt.Print("--------------------- Test: handlerRegister ---------------------\n")

	var cfg = config.Read()
	db, errDB := sql.Open("postgres", cfg.Db_url)
	if errDB != nil {
		fmt.Printf("%v \n", errDB)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	cases := []struct {
		title          string
		pos            int
		state          config.State
		precon         string
		command        config.Cmd
		expectedOutput string
		expectedError  error
	}{
		{
			title: "no username",
			pos:   1,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			precon: "delete_users",
			command: config.Cmd{
				Name:      "register",
				Arguments: []string{},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("invalid number of arguments provided"),
		},
		{
			title: "happy path",
			pos:   2,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			precon: "",
			command: config.Cmd{
				Name: "register",
				Arguments: []string{
					"ashley",
				},
			},
			expectedOutput: "User ashley was created. \n",
			expectedError:  nil,
		},
		{
			title: "Too many arguments",
			pos:   3,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			precon: "",
			command: config.Cmd{
				Name: "register",
				Arguments: []string{
					"login",
					"johnson",
				},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("invalid number of arguments provided"),
		},
	}

	for _, c := range cases {

		fmt.Print("---------------------\n")
		fmt.Printf("%s \n", c.title)
		output, err := testkit.CaptureOutput(func() error {
			return handlerRegister(&c.state, c.command)
		})
		fmt.Printf("Actual output: %v\n", output)
		fmt.Printf("Expected output: %v\n", c.expectedOutput)
		fmt.Printf("Actual error: %v\n", err)
		fmt.Printf("Expected error: %v\n", c.expectedError)
		if err != nil {
			if c.expectedError == nil {
				t.Errorf("expected different error")
			} else if err.Error() != c.expectedError.Error() {
				t.Errorf("expected different error")
			}
		} else if c.expectedError != nil {
			t.Errorf("expected different error")
		}
		if !strings.Contains(output, c.expectedOutput) {
			t.Errorf("expected different output")
		}
		if c.pos == len(cases) {
			ctxt := context.Background()
			c.state.Db_Queries.DeleteAllUsers(ctxt)
		}
	}
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Tobias1006/gatorgo/internal/database"
	"github.com/Tobias1006/gatorgo/internal/testkit"
	"github.com/google/uuid"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func TestHandlerLogin(t *testing.T) {
	fmt.Print("--------------------- Test: handlerLogin ---------------------\n")

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
		command        config.Cmd
		expectedOutput string
		expectedError  error
		expectedConfig config.Config
	}{
		{
			title: "No command, no username",
			pos:   1,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			command: config.Cmd{
				Name:      "",
				Arguments: []string{},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("incorrect number of arguments for the command found"),
			expectedConfig: *cfg,
		},
		{
			title: "happy path",
			pos:   2,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			command: config.Cmd{
				Name: "login",
				Arguments: []string{
					"Ashley",
				},
			},
			expectedOutput: "User Ashley has been set successfully \n",
			expectedError:  nil,
			expectedConfig: config.Config{
				Db_url:            cfg.Db_url,
				Current_user_name: "Ashley",
			},
		},
		{
			title: "Too many arguments",
			pos:   3,
			state: config.State{
				Config:     cfg,
				Db_Queries: dbQueries,
			},
			command: config.Cmd{
				Name: "login",
				Arguments: []string{
					"login",
					"Ashley",
					"Trevor",
				},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("incorrect number of arguments for the command found"),
			expectedConfig: *cfg,
		},
	}

	for _, c := range cases {
		// setting up database for user-readout
		if c.pos == 1 {
			ctxt := context.Background()
			users := []string{"Ashley", "John", "Trevor"}
			for _, name := range users {
				c.state.Db_Queries.CreateUser(ctxt, database.CreateUserParams{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Name:      name,
				})
			}
			c.state.Config.Current_user_name = "Trevor"
		}

		//capturing and assessing output
		fmt.Print("---------------------\n")
		fmt.Printf("%s \n", c.title)
		output, err := testkit.CaptureOutput(func() error {
			return handlerLogin(&c.state, c.command)
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
		if output != c.expectedOutput {
			t.Errorf("expected different output")
		}
		if c.pos == len(cases) {
			ctxt := context.Background()
			c.state.Db_Queries.DeleteAllUsers(ctxt)
		}
	}
}

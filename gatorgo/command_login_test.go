package main

import (
	"fmt"
	"testing"

	"github.com/Tobias1006/gatorgo/internal/testkit"

	"github.com/Tobias1006/gatorgo/internal/config"
)

var cfg = config.Read()

func TestHandlerLogin(t *testing.T) {
	cases := []struct {
		title          string
		state          config.State
		command        config.Cmd
		expectedOutput string
		expectedError  error
		expectedConfig config.Config
	}{
		{
			title: "No command, no username",
			state: config.State{
				Config: cfg,
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
			state: config.State{
				Config: cfg,
			},
			command: config.Cmd{
				Name: "login",
				Arguments: []string{
					"ashley",
				},
			},
			expectedOutput: "User ashley has been set successfully \n",
			expectedError:  nil,
			expectedConfig: config.Config{
				Db_url:            cfg.Db_url,
				Current_user_name: "Ashley",
			},
		},
		{
			title: "Too many arguments",
			state: config.State{
				Config: cfg,
			},
			command: config.Cmd{
				Name: "login",
				Arguments: []string{
					"login",
					"ashley",
					"johnson",
				},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("incorrect number of arguments for the command found"),
			expectedConfig: *cfg,
		},
	}

	for _, c := range cases {
		fmt.Print("---------------------\n")
		fmt.Printf("%s \n", c.title)
		output, err := testkit.CaptureOutput(func() error {
			return handlerLogin(&c.state, c.command)
		})
		fmt.Printf("Actual output: %v\n", output)
		fmt.Printf("Actual error: %v\n", err)
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
	}
}

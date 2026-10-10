package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Tobias1006/gatorgo/internal/testkit"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func TestHandlerRegister(t *testing.T) {

	cases := []struct {
		title          string
		state          config.State
		command        config.Cmd
		expectedOutput string
		expectedError  error
	}{
		{
			title: "no username",
			state: config.State{
				Config: cfg,
			},
			command: config.Cmd{
				Name:      "register",
				Arguments: []string{},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("invalid number of arguments provided"),
		},
		{
			title: "happy path",
			state: config.State{
				Config: cfg,
			},
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
			state: config.State{
				Config: cfg,
			},
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
		{
			title: "Repeated name",
			state: config.State{
				Config: cfg,
			},
			command: config.Cmd{
				Name: "register",
				Arguments: []string{
					"ashley",
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
		if !strings.Contains(output, c.expectedOutput) {
			t.Errorf("expected different output")
		}
	}
}

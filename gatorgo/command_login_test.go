package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func captureOutput(f func() error) (string, error) {
	// save the real stdout so we can restore it later
	old := os.Stdout

	// create an os.Pipe() -- it gives you a connected reader and writer
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f() // run the code that prints to stdout

	w.Close()
	os.Stdout = old // restore it!

	out, _ := io.ReadAll(r)
	return string(out), err
}

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
				Name:      "login",
				Arguments: []string{},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("incorrect number of arguments for the command found"),
			expectedConfig: *cfg,
		},
		{
			title: "No username",
			state: config.State{
				Config: cfg,
			},
			command: config.Cmd{
				Name: "login",
				Arguments: []string{
					"login",
				},
			},
			expectedOutput: "Too few arguments. \n",
			expectedError:  fmt.Errorf("incorrect number of arguments for the command found \n exit status 1"),
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
					"login",
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
		fmt.Printf("%s \n", c.title)
		output, err := captureOutput(func() error {
			return handlerLogin(&c.state, c.command)
		})
		fmt.Printf("Actual output: %v\n", output)
		fmt.Printf("Actual error: %v\n", err)
		if err != nil {
			if err.Error() != c.expectedError.Error() {
				t.Errorf("expected different error")
			} else if output != c.expectedOutput {
				t.Errorf("expected different output")
			}
		}

	}
}

package main

import (
	"fmt"
	"testing"

	"github.com/Tobias1006/gatorgo/internal/config"
)

func TestGetConfigFilePath(t *testing.T) {
	cases := []struct {
		title              string
		expectedResultPath string
		expectedError      error
	}{
		{
			title:              "Config found",
			expectedResultPath: "/home/tobiasg/.gatorconfig.json",
			expectedError:      nil,
		},
	}

	for _, c := range cases {
		fmt.Print("---------------------\n")
		fmt.Printf("%s \n", c.title)
		path, err := config.GetConfigFilePath()
		fmt.Printf("Actual path: %v\n", path)
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
		if path != c.expectedResultPath {
			t.Errorf("expected different path")
		}
	}
}

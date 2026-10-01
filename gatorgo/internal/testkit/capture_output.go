package testkit

import (
	"io"
	"os"
)

func CaptureOutput(f func() error) (string, error) {
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

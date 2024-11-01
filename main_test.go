package main

import (
	"os"
	"testing"
)

func TestMainProgram(t *testing.T) {
	// Create a temporary test file
	content := []byte("...#\n...#\n...#\n...#\n\n..##\n..##\n....\n....\n")
	tmpfile, err := os.CreateTemp("", "test.*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test with valid input
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", tmpfile.Name()}

	main()

	// Test with invalid arguments
	os.Args = []string{"cmd"}
	main()
}

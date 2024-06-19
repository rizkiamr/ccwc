package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCcwc(t *testing.T) {
	// Create a temporary file with known content
	tmpfile, err := ioutil.TempFile("", "example.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up after the test

	content := "Hello, World!"
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Run the ccwc command with the -c flag
	cmd := exec.Command("./ccwc", "-c", tmpfile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// Get the expected output size
	expectedOutput := len(content)
	expected := fmt.Sprintf("%8d %s\n", expectedOutput, tmpfile.Name())

	// Check the output
	if strings.TrimSpace(string(output)) != strings.TrimSpace(expected) {
		t.Errorf("Expected output: %q, but got: %q", expected, string(output))
	}
}

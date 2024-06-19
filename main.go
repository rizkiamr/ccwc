package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define the -c flag
	countBytes := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	// Get the remaining arguments after parsing flags
	args := flag.Args()

	// Check if the -c flag is provided and exactly one filename is given
	if *countBytes && len(args) == 1 {
		filename := args[0]

		// Read the file content
		fileInfo, err := os.Stat(filename)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		// Get the file size in bytes
		fileSize := fileInfo.Size()

		// Print the result in the specified format
		fmt.Printf("%8d %s\n", fileSize, filename)
	} else {
		// Print usage if the flag or filename is missing
		fmt.Println("Usage: ccwc -c <filename>")
	}
}

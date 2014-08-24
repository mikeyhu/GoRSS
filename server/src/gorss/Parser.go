package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Please provide a Url to retrieve\n")
		return
	}

	result, err := LoadUrl(args[0])
	if err != nil {
		// handle error
	}
	for _, story := range result {
		fmt.Printf("Title: %v\n", story.Title)
	}
}

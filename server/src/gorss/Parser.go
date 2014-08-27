package main

import (
	"fmt"
	"gorss/state"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Please provide a Url to retrieve\n")
		return
	}

	stories, err := LoadUrl(args[0])
	if err != nil {
		return
	}
	err = state.IngestStories("localhost:27000", stories)

	if err != nil {
		return
	}
}

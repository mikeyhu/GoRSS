package main

import (
	"fmt"
	"gorss/controllers"
	"gorss/state"
	"os"
)

const CONNECTION = "localhost:27000"
const PORT = "8080"

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Please provide a Url to retrieve\n")
		return
	}

	stories, err := LoadUrl(args[0])
	if err != nil {
		fmt.Printf("Err:%v", err)
		return
	}

	repo := state.GetStoryRepo(CONNECTION)

	repo.Insert(stories)

	if err != nil {
		fmt.Printf("Err:%v", err)
		return
	}

	err = controllers.StartController(CONNECTION, PORT)
	if err != nil {
		fmt.Printf("Err:%v", err)
		return
	}
	fmt.Printf("Hello...")
}

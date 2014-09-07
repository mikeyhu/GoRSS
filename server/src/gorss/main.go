package main

import (
	"gorss/collector"
	"gorss/controllers"
	"gorss/state"
	"log"
	"time"
)

const CONNECTION = "localhost:27000"
const PORT = "8080"

func main() {

	feeds := state.GetFeedRepo(CONNECTION)
	store := state.GetStoryRepo(CONNECTION)

	duration, err := time.ParseDuration("1m")

	go collector.ScheduleStoryCollection(feeds, store, duration)

	err = controllers.StartController(CONNECTION, PORT)
	if err != nil {
		log.Printf("Err:%v", err)
		return
	}
}

package main

import (
	"gorss/collector"
	"gorss/controllers"
	"gorss/state"
	"log"
	"os"
	"time"
)

func main() {

	CONNECTION := envOrDefault("CONNECTION", "localhost:27000")
	PORT := envOrDefault("PORT", "8080")
	DURATION := envOrDefault("DURATION", "15m")

	feeds := state.GetFeedRepo(CONNECTION)
	store := state.GetStoryRepo(CONNECTION)

	duration, err := time.ParseDuration(DURATION)

	go collector.ScheduleStoryCollection(feeds, store, duration)

	log.Printf("Starting server on port %v", PORT)
	err = controllers.StartController(CONNECTION, PORT)
	if err != nil {
		log.Printf("Err:%v", err)
		return
	}
}

func envOrDefault(envName string, defaultValue string) (result string) {
	result = os.Getenv(envName)
	if result == "" {
		result = defaultValue
	}
	return result

}

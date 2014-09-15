package controllers

import (
	"github.com/gorilla/mux"
	"gorss/state"
	"net/http"
)

var (
	storyRepo state.StoryRepo
	feedRepo  state.FeedRepo
)

func StartController(connection string, port string) error {
	storyRepo = state.GetStoryRepo(connection)
	feedRepo = state.GetFeedRepo(connection)

	r := mux.NewRouter()

	//Static Handlers
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("../client/dist/"))))

	http.Handle("/", r)

	http.Handle("/feeds/", FeedsController())
	http.Handle("/stories/", StoriesController())

	return http.ListenAndServe(":"+port, nil)
}

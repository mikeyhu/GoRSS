package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorss/state"
	"net/http"
)

var (
	storyRepo state.StoryRepo
	feedRepo  state.FeedRepo
)

func LatestStoriesHandler(w http.ResponseWriter, r *http.Request) {
	stories, _ := storyRepo.All()

	w.Header().Set("Content-Type", "application/json")

	if len(stories) > 0 {
		js, err := json.Marshal(stories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	} else {
		w.Write([]byte("[]"))
	}
}

func allFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := feedRepo.All()

	js, err := json.Marshal(feeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartController(connection string, port string) error {
	storyRepo = state.GetStoryRepo(connection)
	feedRepo = state.GetFeedRepo(connection)

	r := mux.NewRouter()
	//Story Handlers
	r.HandleFunc("/stories/latest", LatestStoriesHandler)

	//Feed Handlers
	r.HandleFunc("/feeds/all", allFeedsHandler)

	//Static Handlers

	http.Handle("/", r)
	return http.ListenAndServe(":"+port, nil)
}

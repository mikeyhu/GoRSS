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

func StartController(connection string, port string) error {
	storyRepo = state.GetStoryRepo(connection)
	feedRepo = state.GetFeedRepo(connection)

	r := mux.NewRouter()
	//Story Handlers
	r.HandleFunc("/stories/latest", LatestStoriesHandler)

	//Static Handlers
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("../client/dist/"))))

	http.Handle("/", r)

	http.Handle("/feeds/", FeedsController())

	return http.ListenAndServe(":"+port, nil)
}

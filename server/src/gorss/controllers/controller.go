package controllers

import (
	"net/http"
	"gorss/state"
	"encoding/json"
)

var storyRepo state.StoryRepo

func getLatestStories(w http.ResponseWriter, r *http.Request) {
	stories, err := storyRepo.All()

	js, err := json.Marshal(stories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartController(repo state.StoryRepo, port string) error {
	storyRepo = repo
	http.HandleFunc("/stories/latest", getLatestStories)
	return http.ListenAndServe(":" + port, nil)
}


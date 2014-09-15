package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, _ := storyRepo.Tags()

	w.Header().Set("Content-Type", "application/json")

	if len(tags) > 0 {
		js, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	} else {
		w.Write([]byte("[]"))
	}
}

func StoriesController() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/stories/latest", LatestStoriesHandler)
	r.HandleFunc("/stories/tags", TagsHandler)
	return
}

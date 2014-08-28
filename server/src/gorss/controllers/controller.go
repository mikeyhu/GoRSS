package controllers

import (
	"net/http"
	"gorss/state"
	"encoding/json"
)

var connection string

func getStories(w http.ResponseWriter, r *http.Request) {

	stories, err := state.GetStories(connection)

	js, err := json.Marshal(stories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StartController(mongoConnection string, port string) error {
	connection = mongoConnection
	http.HandleFunc("/stories", getStories)
	return http.ListenAndServe(":" + port, nil)

}


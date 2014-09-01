package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorss/domain"
	"net/http"
)

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

func insertFeedHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var feed domain.Feed
	err := decoder.Decode(&feed)
	if err != nil {
		http.Error(w, "Invalid Request", 400)
	} else {
		feedRepo.Insert(feed)
	}
}

func FeedsController() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/feeds/", insertFeedHandler).Methods("POST")
	r.HandleFunc("/feeds/all", allFeedsHandler)
	return
}

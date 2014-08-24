package main

import (
	"gorss/atom"
	"gorss/domain"
	"gorss/rss"
	"io/ioutil"
	"net/http"
)

func LoadUrl(url string) (result []domain.Story, err error) {
	resp, err := http.Get("http://feeds.bbci.co.uk/news/video_and_audio/uk/rss.xml")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return LoadFeed(string(body))
}

func LoadFeed(data string) (result []domain.Story, err error) {
	result, err = atom.LoadStories(data)
	if err == nil {
		return
	}
	return rss.LoadStories(data)
}

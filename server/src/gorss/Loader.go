package main

import (
	"gorss/atom"
	"gorss/domain"
	"gorss/rss"
	"io/ioutil"
	"net/http"
)

func LoadUrl(url string) (result []domain.Story, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return LoadFeed(string(body))
}

func LoadFeed(data string) (result []domain.Story, err error) {
	if result, err = atom.LoadStories(data); err == nil {
		return
	}
	return rss.LoadStories(data)
}

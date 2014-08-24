package main

import (
	"gorss/atom"
	"gorss/domain"
	"gorss/rss"
)

func LoadFeed(data string) (result []domain.Story, err error) {
	result, err = atom.LoadStories(data)
	if err == nil {
		return
	}
	result, err = rss.LoadStories(data)
	return
}

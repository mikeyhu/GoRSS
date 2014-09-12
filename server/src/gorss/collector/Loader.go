package collector

import (
	"gorss/atom"
	"gorss/domain"
	"gorss/rss"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadUrl(url string, tags []string) (result []domain.Story, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result, err = loadFeed(string(body))
	if err != nil {
		return
	}
	log.Printf("Retrieved %v stories from url: %v", len(result), url)
	result = updateStories(url, tags, result)

	return
}

func loadFeed(data string) (result []domain.Story, err error) {
	if result, err = atom.LoadStories(data); err == nil {
		return
	}
	return rss.LoadStories(data)
}

func updateStories(url string, tags []string, stories []domain.Story) []domain.Story {
	results := make([]domain.Story, len(stories))
	for i, v := range stories {
		results[i] = v
		results[i].Url = url
		results[i].Tags = tags
	}
	return results
}

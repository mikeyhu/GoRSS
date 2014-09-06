package collector

import (
	"gorss/atom"
	"gorss/domain"
	"gorss/rss"
	"io/ioutil"
	"log"
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
	result, err = loadFeed(string(body))

	log.Printf("Retrieved %v stories from url: %v", len(result), url)

	return
}

func loadFeed(data string) (result []domain.Story, err error) {
	if result, err = atom.LoadStories(data); err == nil {
		return
	}
	return rss.LoadStories(data)
}

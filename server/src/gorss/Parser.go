package main

import (
	"fmt"
	"gorss/rss"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://feeds.bbci.co.uk/news/video_and_audio/uk/rss.xml")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	rssBody := string(body)
	//println(rssBody)
	parsed, _ := rss.Parse(rssBody)

	for _, story := range rss.Normalise(parsed) {
		fmt.Printf("Title: %v\n", story.Title)
	}

}

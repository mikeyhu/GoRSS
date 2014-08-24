package main

import (
	"fmt"
)

func main() {
	result, err := LoadUrl("http://feeds.bbci.co.uk/news/video_and_audio/uk/rss.xml")
	if err != nil {
		// handle error
	}
	for _, story := range result {
		fmt.Printf("Title: %v\n", story.Title)
	}
}

package rss

import (
	"encoding/xml"
	"fmt"
	"gorss/data"
)

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func Parse(data string) (result Rss, err error) {
	result = Rss{}
	err = xml.Unmarshal([]byte(data), &result)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return
}

func Normalise(parsedData Rss) []data.Story {
	var results = make([]data.Story, len(parsedData.Channel.Items))

	for pos, element := range parsedData.Channel.Items {
		results[pos] = data.Story{
			Title: element.Title,
			Link:  element.Link}
	}
	return results
}

package rss

import (
	"encoding/xml"
	"errors"
	"gorss/domain"
	"time"
)

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Guid    string `xml:"guid"`
	Pubdate string `xml:"pubDate"`
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
	return
}

func Normalise(parsedData Rss) []domain.Story {
	var results = make([]domain.Story, len(parsedData.Channel.Items))

	for pos, element := range parsedData.Channel.Items {
		date, _ := time.Parse(time.RFC1123, element.Pubdate)
		results[pos] = domain.Story{
			Title: element.Title,
			Link:  element.Link,
			Id:    element.Guid,
			Date:  date,
			State: "new"}
	}
	return results
}

func LoadStories(data string) (stories []domain.Story, err error) {
	var result Rss
	result, err = Parse(data)
	if err != nil {
		return
	}
	stories = Normalise(result)
	if len(stories) == 0 {
		err = errors.New("No stories found")
	}
	return
}

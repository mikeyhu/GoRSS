package atom

import (
	"encoding/xml"
	"errors"
	"gorss/domain"
)

type Entry struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Feed struct {
	Title   string  `xml:"title"`
	Entries []Entry `xml:"entry"`
}

func Parse(data string) (rss Feed, err error) {
	rss = Feed{}
	err = xml.Unmarshal([]byte(data), &rss)
	return
}

func Normalise(parsedData Feed) []domain.Story {
	var results = make([]domain.Story, len(parsedData.Entries))

	for pos, element := range parsedData.Entries {
		results[pos] = domain.Story{
			Title: element.Title,
			Link:  element.Link}
	}
	return results
}

func LoadStories(data string) (stories []domain.Story, err error) {
	var result Feed
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

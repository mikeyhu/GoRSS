package atom

import (
	"encoding/xml"
	"gorss/data"
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

func Normalise(parsedData Feed) []data.Story {
	var results = make([]data.Story, len(parsedData.Entries))

	for pos, element := range parsedData.Entries {
		results[pos] = data.Story{
			Title: element.Title,
			Link:  element.Link}
	}
	return results
}

func LoadStories(data string) (stories []data.Story, err error) {
	var result Feed
	result, err = Parse(data)
	if err != nil {
		return
	}
	stories = Normalise(result)
	return
}

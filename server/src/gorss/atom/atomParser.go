package atom

import (
	"encoding/xml"
	"errors"
	"gorss/domain"
	"time"
)

type Entry struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Updated string `xml:"updated"`
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

func Normalise(parsedData Feed) (results []domain.Story) {
	results = make([]domain.Story, len(parsedData.Entries))

	for pos, element := range parsedData.Entries {
		date, _ := time.Parse(time.RFC1123, element.Updated)
		results[pos] = domain.Story{
			Title: element.Title,
			Link:  element.Link,
			Date:  date}
	}
	return
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

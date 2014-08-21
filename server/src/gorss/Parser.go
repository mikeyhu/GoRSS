package gorss

import (
	"gorss/atom"
)

type Story struct {
	Title string
	Link  string
}

func NormaliseAtom(parsedData atom.Feed) []Story {
	var results = make([]Story, len(parsedData.Entries))

	for pos, element := range parsedData.Entries {
		results[pos] = Story{
			Title: element.Title,
			Link:  element.Link}
	}
	return results
}

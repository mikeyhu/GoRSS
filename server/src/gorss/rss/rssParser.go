package rss

import (
	"encoding/xml"
	"fmt"
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

func Parse(data string) (rss Rss, err error) {
	rss = Rss{}
	err = xml.Unmarshal([]byte(data), &rss)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return

}

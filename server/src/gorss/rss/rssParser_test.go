package rss

import (
	"testing"
)

var testData = `
<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:taxo="http://purl.org/rss/1.0/modules/taxonomy/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:media="http://search.yahoo.com/mrss/" version="2.0">
  <channel>
	<title>Sport | The Guardian</title>
	<link>http://www.theguardian.com/uk/sport</link>
	<description>Sport news, results, fixtures, blogs and comments on UK and world sport from the Guardian, the world's leading liberal voice</description>
	<language>en-gb</language>
	<copyright>Guardian News and Media Limited or its affiliated companies. All rights reserved. 2014</copyright>
	<pubDate>Tue, 24 Jun 2014 20:02:55 GMT</pubDate>
	<lastBuildDate>Tue, 24 Jun 2014 20:02:55 GMT</lastBuildDate>
	<ttl>5</ttl>
	<dc:date>2014-06-24T20:02:55Z</dc:date>
	<dc:language>en-gb</dc:language>
	<dc:rights>Guardian News and Media Limited or its affiliated companies. All rights reserved. 2014</dc:rights>
	<item>
	  <title>Japan v Colombia: World Cup 2014 live!</title>
	  <link>http://feeds.theguardian.com/c/34708/f/666716/s/3bd61efc/sc</link>
	</item>
	<item>
	  <title>Germany v Brazil: World Cup 2014 live!</title>
	  <link>http://feeds.theguardian.com/c/34708</link>
	</item>
  </channel>
</rss>
`

func TestParseSuccess(t *testing.T) {
	var result, err = Parse(testData)

	if err != nil {
		t.Errorf("Parse() returned %v", err)
	}
	expectedTitle := "Sport | The Guardian"
	if result.Channel.Title != expectedTitle {
		t.Errorf("Parse() result.Title = '%v' wanted '%v' ", result.Channel.Title, expectedTitle)
	}
}

func TestParseFailInvalidXML(t *testing.T) {
	testData := "<bobbob><bil>"

	_, err := Parse(testData)

	if err == nil {
		t.Errorf("Parse() returned %v", err)
	}
}

func TestNormalise(t *testing.T) {
	var parsed, _ = Parse(testData)
	var result = Normalise(parsed)

	if len(result) != 2 {
		t.Errorf("Normalise() returned %v", len(result))
	}
}

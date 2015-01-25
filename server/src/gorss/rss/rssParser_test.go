package rss

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type RssParserSuite struct{}

var _ = Suite(&RssParserSuite{})

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
	  <guid isPermaLink="false">http://feeds.theguardian.com/c/34708/f/666716/s/3bd61efc/sc</guid>
	  <pubDate>Tue, 24 Jun 2014 20:02:55 GMT</pubDate>
	</item>
	<item>
	  <title>Germany v Brazil: World Cup 2014 live!</title>
	  <link>http://feeds.theguardian.com/c/34708</link>
	  <pubDate>Tue, 24 Jun 2014 20:02:55 GMT</pubDate>
	</item>
  </channel>
</rss>
`

func (s *RssParserSuite) TestParseSuccess(c *C) {
	var result, err = Parse(testData)

	c.Assert(err, IsNil)
	c.Assert(result.Channel.Title, Equals, "Sport | The Guardian")
}

func (s *RssParserSuite) TestParseFailInvalidXML(c *C) {
	testData := "<bobbob><bil>"

	_, err := Parse(testData)

	c.Assert(err, Not(IsNil))
}

func (s *RssParserSuite) TestNormalise(c *C) {
	var parsed, _ = Parse(testData)
	var result = Normalise(parsed)

	c.Assert(result, HasLen, 2)
	c.Assert(result[0].Id, Equals, "http://feeds.theguardian.com/c/34708/f/666716/s/3bd61efc/sc")

	expectedDate, err := time.Parse(time.RFC1123, "Tue, 24 Jun 2014 20:02:55 GMT")
	c.Assert(err, IsNil)
	c.Assert(result[0].Date, Equals, expectedDate)
}

func (s *RssParserSuite) TestNormaliseAddsDefaultValues(c *C) {
	var parsed, _ = Parse(testData)
	var result = Normalise(parsed)

	c.Assert(result[0].State, Equals, "new")
}

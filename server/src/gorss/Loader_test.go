package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type LoaderSuite struct{}

var _ = Suite(&LoaderSuite{})

var testAtom = `
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Example Feed</title>
	<entry>
		<title>Atom-Powered Robots Run Amok</title>
		<link href="http://example.org/2003/12/13/atom03" />
		<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
	</entry>
</feed>
`

var testRss = `
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
	<item>
	  <title>Japan v Colombia: World Cup 2014 live!</title>
	  <link>http://feeds.theguardian.com/c/34708/f/666716/s/3bd61efc/sc</link>
	  <guid isPermaLink="false">http://feeds.theguardian.com/c/34708/f/666716/s/3bd61efc/sc</guid>
	  <pubDate>Tue, 24 Jun 2014 20:02:55 GMT</pubDate>
	</item>
  </channel>
</rss>
`

func (s *LoaderSuite) TestLoadRss(c *C) {
	result, err := LoadFeed(testRss)
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, 1)
}

func (s *LoaderSuite) TestLoadAtom(c *C) {
	result, err := LoadFeed(testAtom)
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, 1)
}

func (s *LoaderSuite) TestInvalid(c *C) {
	_, err := LoadFeed("<a></a>")
	c.Assert(err, Not(IsNil))
}

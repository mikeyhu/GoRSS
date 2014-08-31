package atom

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type AtomParserSuite struct{}

var _ = Suite(&AtomParserSuite{})

var testData = `
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Example Feed</title>
	<subtitle>A subtitle.</subtitle>
	<link href="http://example.org/feed/" rel="self" />
	<link href="http://example.org/" />
	<id>urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6</id>
	<updated>2003-12-13T18:30:02Z</updated>
	<entry>
		<title>Atom-Powered Robots Run Amok</title>
		<link href="http://example.org/2003/12/13/atom03" />
		<link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
		<link rel="edit" href="http://example.org/2003/12/13/atom03/edit"/>
		<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
		<updated>2003-12-13T18:30:02Z</updated>
		<summary>Some text.</summary>
		<content type="xhtml">
		   <div xmlns="http://www.w3.org/1999/xhtml">
			  <p>This is the entry content.</p>
		   </div>
		</content>
		<author>
			<name>John Doe</name>
			<email>johndoe@example.com</email>
	   </author>
	</entry>
</feed>
`

var testEmptyData = `
<feed xmlns="http://www.w3.org/2005/Atom">
	<title>Example Feed</title>
	<subtitle>A subtitle.</subtitle>
	<link href="http://example.org/feed/" rel="self" />
	<link href="http://example.org/" />
	<id>urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6</id>
	<updated>2003-12-13T18:30:02Z</updated>
</feed>
`

func (s *AtomParserSuite) TestParseSuccess(c *C) {
	var result, err = Parse(testData)

	c.Assert(err, IsNil)
	c.Assert(result.Title, Equals, "Example Feed")
}

func (s *AtomParserSuite) TestParseFailInvalidXML(c *C) {
	testData := "<bobbob><bil>"
	_, err := Parse(testData)

	c.Assert(err, Not(IsNil))
}

func (s *AtomParserSuite) TestNormalise(c *C) {
	var parsed, _ = Parse(testData)
	var result = Normalise(parsed)

	c.Assert(result, HasLen, 1)
	c.Assert(result[0].Title, Equals, "Atom-Powered Robots Run Amok")
}

func (s *AtomParserSuite) TestLoadStories(c *C) {
	var result, err = LoadStories(testData)

	expectedDate, _ := time.Parse(time.RFC1123, "2003-12-13T18:30:02Z")
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, 1)
	c.Assert(result[0].Title, Equals, "Atom-Powered Robots Run Amok")
	c.Assert(result[0].Date, Equals, expectedDate)
}

func (s *AtomParserSuite) TestLoadEmptyStories(c *C) {
	var result, err = LoadStories(testEmptyData)

	c.Assert(err, Not(IsNil))
	c.Assert(result, HasLen, 0)
}

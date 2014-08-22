package atom

import "testing"

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

func TestParseSuccess(t *testing.T) {
	var result, err = Parse(testData)

	if err != nil {
		t.Errorf("Parse() returned %v", err)
	}
	expectedTitle := "Example Feed"
	if result.Title != expectedTitle {
		t.Errorf("Parse() result.Title = %v wanted %v ", result.Title, expectedTitle)
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

	if len(result) != 1 {
		t.Errorf("NormaliseAtom() returned %v", len(result))
	}
	if result[0].Title != "Atom-Powered Robots Run Amok" {
		t.Errorf("NormaliseAtom() returned title %v", result[0].Title)
	}
}

func TestLoadStories(t *testing.T) {
	var result, err = LoadStories(testData)

	if err != nil {
		t.Errorf("LoadStories() returned err")
	}
	if len(result) != 1 {
		t.Errorf("LoadStories() returned %v", len(result))
	}
	if result[0].Title != "Atom-Powered Robots Run Amok" {
		t.Errorf("LoadStories() returned title %v", result[0].Title)
	}
}

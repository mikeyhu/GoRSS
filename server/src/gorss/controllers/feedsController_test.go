package controllers

import (
	"bytes"
	. "gopkg.in/check.v1"
	"gorss/state"
	"net/http"
	"net/http/httptest"
)

type FeedControllerSuite struct{}

var _ = Suite(&FeedControllerSuite{})

func (s *FeedControllerSuite) SetUpTest(c *C) {
	feedRepo = state.GetFeedRepo(CONNECTION)
	feedRepo.Clear()
}

func (s *FeedControllerSuite) TestInsertFeed(c *C) {
	//Given
	resp := httptest.NewRecorder()
	uri := "/feeds/"

	toInsert := `{"url" : "http://abc/", "Tags" : ["Tech", "Programming"]}`
	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(toInsert))
	if err != nil {
		c.Fatal(err)
	}
	insertFeedHandler(resp, req)

	c.Assert(resp.Code, Equals, 200)

	inserted, err := feedRepo.All()
	c.Assert(inserted, HasLen, 1)
}

func (s *FeedControllerSuite) TestUpdateFeed(c *C) {
	//Given
	resp := httptest.NewRecorder()
	uri := "/feeds/"

	toInsert := `{"url" : "http://abc/", "Tags" : ["Tech", "Programming"]}`

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(toInsert))
	if err != nil {
		c.Fatal(err)
	}
	insertFeedHandler(resp, req)

	inserted, err := feedRepo.All()

	id := inserted[0].Id.Hex()

	toUpdate := string("{\"id\" : \"" + id + "\", \"url\" : \"http://xyz/\", \"Tags\" : [\"Tech\", \"Programming\"]}")

	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", string("/feeds/"+id), bytes.NewBufferString(toUpdate))
	updateFeedHandler(resp, req)

	inserted, err = feedRepo.All()

	c.Assert(inserted, HasLen, 1)
	c.Assert(inserted[0].Url, Equals, "http://xyz/")
}

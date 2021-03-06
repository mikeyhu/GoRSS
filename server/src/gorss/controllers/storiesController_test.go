package controllers

import (
	. "gopkg.in/check.v1"
	"gorss/domain"
	"net/http"
	"net/http/httptest"
)

func (s *ControllerSuite) TestLatestStoriesHandler_ReturnsEmpty(c *C) {
	//Given
	resp := httptest.NewRecorder()
	uri := "/stories/latest"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		c.Fatal(err)
	}
	//When
	LatestStoriesHandler(resp, req)

	//Then
	c.Assert(resp.Code, Equals, http.StatusOK)
	c.Assert(resp.Body.String(), Equals, "[]")
}

func (s *ControllerSuite) TestLatestStoriesHandler_ReturnsAnItem(c *C) {
	//Given
	resp := httptest.NewRecorder()
	uri := "/stories/latest"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		c.Fatal(err)
	}

	var testStory = domain.Story{
		Title: "Another story",
		Id:    "another_story",
		State: "new"}
	stories := []domain.Story{testStory}
	storyRepo.Insert(stories)

	//When
	LatestStoriesHandler(resp, req)

	//Then
	c.Assert(resp.Code, Equals, http.StatusOK)
	c.Assert(
		resp.Body.String(),
		Equals,
		`[{"Title":"Another story","Link":"","Id":"another_story","Date":"0001-01-01T00:00:00Z","Url":"","Tags":[],"State":"new"}]`)
}

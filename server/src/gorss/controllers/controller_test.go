package controllers

import (
	"gorss/domain"
	"gorss/state"
	"net/http"
	"net/http/httptest"
	"testing"
)

var CONNECTION = "localhost:27000"

func TestLatestStoriesShouldReturnEmptyArray(t *testing.T) {

	resp := httptest.NewRecorder()

	uri := "/stories/latest"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	storyRepo = state.GetStoryRepo(CONNECTION)
	storyRepo.Clear()

	LatestStoriesHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Response status expected: %v:\n\t recieved: %v", "200", resp.Code)
	}

	body := resp.Body.String()
	if body != "[]" {
		t.Errorf("Response body expected :[] \n\t recieved: %v", body)
	}
}

func TestLatestStoriesShouldReturnSingleStory(t *testing.T) {

	resp := httptest.NewRecorder()

	uri := "/stories/latest"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	storyRepo = state.GetStoryRepo(CONNECTION)
	storyRepo.Clear()

	var testStory = domain.Story{
		Title: "Another story",
		Id:    "another_story"}

	stories := []domain.Story{
		testStory}
	storyRepo.Insert(stories)

	LatestStoriesHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Response status expected: %v:\n\t recieved: %v", "200", resp.Code)
	}

	body := resp.Body.String()
	expected := `[{"Title":"Another story","Link":"","Id":"another_story","Date":"0001-01-01T00:00:00Z"}]`
	if body != expected {
		t.Errorf("Response body \n\t expected:%v \n\t recieved:%v", expected, body)
	}
}

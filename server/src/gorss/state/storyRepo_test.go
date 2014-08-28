package state

import (
	"gorss/domain"
	"testing"
	"time"
)

var expectedDate = time.Now()

var testStory1 = domain.Story{
	Title: "A story",
	Id:    "a_story",
	Date:  expectedDate}

var testStory2 = domain.Story{
	Title: "Another story",
	Id:    "another_story",
	Date:  expectedDate}

func TestStoryRepo_Insert(t *testing.T) {
	clearCollection(COLLECTION_STORIES)
	stories := []domain.Story{
		testStory1,
		testStory2}

	repo := GetStoryRepo(CONNECTION)

	err := repo.Insert(stories)

	if err != nil {
		t.Errorf("StoryRepo.Insert() returned %v", err)
	}
}

func TestStoryRepo_All(t *testing.T) {
	clearCollection(COLLECTION_STORIES)
	stories := []domain.Story{
		testStory1,
		testStory2}

	repo := GetStoryRepo(CONNECTION)

	err := repo.Insert(stories)

	result, err := repo.All()

	if err != nil {
		t.Errorf("StoryRepo.All() returned %v", err)
	}
	if len(result) != 2 {
		t.Errorf("StoryRepo.All() returned %v items: \n%v", len(result), result)
	}
}

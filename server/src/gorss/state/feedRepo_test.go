package state

import (
	"gorss/domain"
	"testing"
)

var testFeed1 = domain.Feed{
	Url:  "http://localhost:12345/rss.xml",
	Tags: []string{"News", "Technology"}}

var testFeed2 = domain.Feed{
	Url:  "http://localhost:54321/rss.xml",
	Tags: []string{"News", "Technology"}}

func TestFeedRepo_Insert(t *testing.T) {
	clearCollection(COLLECTION_FEEDS)
	feeds := []domain.Feed{
		testFeed1,
		testFeed2}

	repo := GetFeedRepo(CONNECTION)

	err := repo.Insert(feeds)

	if err != nil {
		t.Errorf("FeedRepo.Insert() returned %v", err)
	}
}

func TestFeedRepo_All(t *testing.T) {
	clearCollection(COLLECTION_FEEDS)
	feeds := []domain.Feed{
		testFeed1,
		testFeed2}

	repo := GetFeedRepo(CONNECTION)

	err := repo.Insert(feeds)

	result, err := repo.All()

	if err != nil {
		t.Errorf("FeedRepo.All() returned %v", err)
	}
	if len(result) != 2 {
		t.Errorf("FeedRepo.All() returned %v", result)
	}
}

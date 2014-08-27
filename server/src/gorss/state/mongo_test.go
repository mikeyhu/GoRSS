package state

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorss/domain"
	"testing"
	"time"
)

var connection = "localhost:27000"

var expectedDate = time.Now()

var testStory1 = domain.Story{
	Title: "A story",
	Id:    "a_story",
	Date:  expectedDate}

var testStory2 = domain.Story{
	Title: "Another story",
	Id:    "another_story",
	Date:  expectedDate}

var testFeed1 = domain.Feed{
	Url:  "http://localhost:12345/rss.xml",
	Tags: []string{"News", "Technology"}}

var testFeed2 = domain.Feed{
	Url:  "http://localhost:54321/rss.xml",
	Tags: []string{"News", "Technology"}}

func TestInsertFeed(t *testing.T) {
	clearCollection(COLLECTION_FEEDS)
	feeds := []domain.Feed{
		testFeed1,
		testFeed2}

	erro := IngestFeeds(connection, feeds)
	if erro != nil {
		t.Errorf("Ingest() returned %v", erro)
	}
}

func TestGetFeed(t *testing.T) {
	clearCollection(COLLECTION_FEEDS)
	feeds := []domain.Feed{
		testFeed1,
		testFeed2}

	_ = IngestFeeds(connection, feeds)

	result, err := GetFeeds(connection)
	if err != nil {
		t.Errorf("GetFeeds() returned %v", err)
	}
	if len(result) != 2 {
		t.Errorf("GetFeeds() returned %v", result)
	}
}

func TestIngestion(t *testing.T) {

	stories := []domain.Story{
		testStory1,
		testStory2}

	erro := IngestStories(connection, stories)
	if erro != nil {
		t.Errorf("Ingest() returned %v", erro)
	}

	session, err := mgo.Dial(connection)
	defer session.Close()
	c := session.DB(DB_NAME).C(COLLECTION_STORIES)

	var result domain.Story

	err = c.Find(bson.M{"title": "A story"}).One(&result)
	if err != nil {
		t.Errorf("Cannot find story: %v", err)
	}
	if result.Title != "A story" {
		t.Errorf("Story.Title: %v", result.Title)
	}
}

func clearCollection(collection string) {
	session, _ := mgo.Dial(connection)

	defer session.Close()

	c := session.DB(DB_NAME).C(collection)
	c.DropCollection()

}

package ingestion

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"gorss/domain"
	"testing"
)

var connection = "localhost:27000"

var testStory1 = domain.Story{
	Title : "A story",
	Id    : "a_story"}

var testStory2 = domain.Story{
	Title : "Another story",
	Id    : "another_story"}

func TestIngestion(t *testing.T) {

	stories := []domain.Story{
		testStory1,
		testStory2}

	erro := Ingest(connection, stories)
	if erro != nil {
		t.Errorf("Ingest() returned %v", erro)
	}


	session, err := mgo.Dial(connection)
	defer session.Close()
	c := session.DB("test").C("stories")

	var result domain.Story

	err = c.Find(bson.M{"title": "A story"}).One(&result)
	if err != nil {
		t.Errorf("Cannot find story: %v", err)
	}


}

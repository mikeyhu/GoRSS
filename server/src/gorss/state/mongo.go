package state

import (
	"gopkg.in/mgo.v2"
	"gorss/domain"
)

const DB_NAME = "RSS"
const COLLECTION_STORIES = "stories"
const COLLECTION_FEEDS = "feeds"

func IngestStories(connection string, stories []domain.Story) (err error) {
	items := make([]interface{}, len(stories))
	for i, v := range stories {
		items[i] = v
	}
	return ingest(connection, COLLECTION_STORIES, items)
}

func IngestFeeds(connection string, feeds []domain.Feed) (err error) {
	items := make([]interface{}, len(feeds))
	for i, v := range feeds {
		items[i] = v
	}
	return ingest(connection, COLLECTION_FEEDS, items)
}

func ingest(connection string, collection string, items []interface{}) (err error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return
	}
	defer session.Close()

	c := session.DB(DB_NAME).C(collection)
	for _, item := range items {
		err = c.Insert(item)
		if err != nil {
			break
		}
	}
	return
}

package state

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorss/domain"
	"log"
)

type FeedRepo struct {
	Collection *mgo.Collection
}

func (r FeedRepo) All() (feeds []domain.Feed, err error) {
	err = r.Collection.Find(bson.M{}).All(&feeds)
	return
}

func (r FeedRepo) Insert(feed domain.Feed) (err error) {

	err = r.Collection.Insert(feed)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Inserted Feed: %v\n", feed.Url)
	}
	return
}

func (r FeedRepo) Clear() (err error) {
	return r.Collection.DropCollection()
}

func (r *FeedRepo) SetMongoCollection(Collection *mgo.Collection) {
	r.Collection = Collection
}

func GetFeedRepo(connection string) FeedRepo {
	var repo = &FeedRepo{}
	getRepo(connection, COLLECTION_FEEDS, repo)
	return *repo
}

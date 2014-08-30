package state

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorss/domain"
)

type FeedRepo struct {
	Collection *mgo.Collection
}

func (r FeedRepo) All() (feeds []domain.Feed, err error) {
	err = r.Collection.Find(bson.M{}).All(&feeds)
	return
}

func (r FeedRepo) Insert(feeds []domain.Feed) (err error) {
	items := make([]interface{}, len(feeds))
	for i, v := range feeds {
		items[i] = v
	}
	for _, item := range items {
		err = r.Collection.Insert(item)
		if err != nil {
			break
		}
	}
	return
}

func (r *FeedRepo) SetMongoCollection(Collection *mgo.Collection) {
	r.Collection = Collection
}

func GetFeedRepo(connection string) FeedRepo {
	var repo = &FeedRepo{}
	getRepo(connection, COLLECTION_FEEDS, repo)
	return *repo
}

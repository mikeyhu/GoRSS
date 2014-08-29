package state

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorss/domain"
)

type (
	StoryRepo struct {
		Collection *mgo.Collection
	}
)

func (r StoryRepo) All() (feeds []domain.Story, err error) {
	err = r.Collection.Find(bson.M{}).All(&feeds)
	return
}

func (r StoryRepo) Insert(feeds []domain.Story) (err error) {
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

func (r *StoryRepo) SetMongoCollection(Collection *mgo.Collection) {
	r.Collection = Collection
}

func GetStoryRepo(connection string) StoryRepo {
	var repo = &StoryRepo{}
	getRepo(connection, COLLECTION_STORIES, repo)
	return *repo
}

package state

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gorss/domain"
	"log"
)

type StoryRepo struct {
	Collection *mgo.Collection
}

func (r StoryRepo) All() (feeds []domain.Story, err error) {
	err = r.Collection.Find(bson.M{}).All(&feeds)
	return
}

func (r StoryRepo) Insert(stories []domain.Story) (err error) {
	items := make([]interface{}, len(stories))
	for i, v := range stories {
		items[i] = v
	}
	for _, item := range items {
		err = r.Collection.Insert(item)
		if err != nil {
			break
		}
	}
	log.Printf("Inserted %v stories", len(stories))
	return
}

func (r StoryRepo) Clear() (err error) {
	return r.Collection.DropCollection()
}

func (r *StoryRepo) SetMongoCollection(Collection *mgo.Collection) {
	r.Collection = Collection
}

func GetStoryRepo(connection string) StoryRepo {
	var repo = &StoryRepo{}
	getRepo(connection, COLLECTION_STORIES, repo)
	return *repo
}

package state

import "gopkg.in/mgo.v2"

const DB_NAME = "RSS"
const COLLECTION_STORIES = "stories"
const COLLECTION_FEEDS = "feeds"

type MongoRepo interface {
	SetMongoCollection(Collection *mgo.Collection)
}

func getRepo(connection string, collection string, repo MongoRepo) {
	var (
		mongoSession *mgo.Session
		err          error
	)
	if mongoSession, err = mgo.Dial(connection); err != nil {
		panic(err)
	}

	database := mongoSession.DB(DB_NAME)
	repo.SetMongoCollection(database.C(collection))
}

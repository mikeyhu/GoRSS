package state

import (
	"gopkg.in/mgo.v2"
)

var CONNECTION = "localhost:27000"

func clearCollection(collection string) {
	session, _ := mgo.Dial(CONNECTION)

	defer session.Close()

	c := session.DB(DB_NAME).C(collection)
	c.DropCollection()

}

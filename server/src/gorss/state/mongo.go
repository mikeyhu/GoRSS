package state

import (
	"gopkg.in/mgo.v2"
	"gorss/domain"
)

func Ingest(connection string, stories []domain.Story) (err error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return
	}
	defer session.Close()

	c := session.DB("test").C("stories")
	for _, story := range stories {
		err = c.Insert(story)
		if err != nil {
			break
		}
	}
	return
}

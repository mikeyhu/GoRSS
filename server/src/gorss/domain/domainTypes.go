package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Story struct {
	Title string
	Link  string
	Id    string
	Date  time.Time
}

type Feed struct {
	Id   bson.ObjectId `json:"id"           bson:"_id"`
	Url  string
	Tags []string
}

package domain

import "time"

type Story struct {
	Title string
	Link  string
	Id    string
	Date  time.Time
}

type Feed struct {
	Url  string
	Tags []string
}

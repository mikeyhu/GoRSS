package state

import (
	. "gopkg.in/check.v1"
	"gorss/domain"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type FeedRepoSuite struct{}

var _ = Suite(&FeedRepoSuite{})

func (s *FeedRepoSuite) SetUpTest(c *C) {
	clearCollection(COLLECTION_FEEDS)
}

var testFeed = domain.Feed{
	Url: "http://localhost:12345/rss.xml", Tags: []string{"News", "Technology"}}

func (s *FeedRepoSuite) TestFeedRepo_Insert(c *C) {
	//Given
	repo := GetFeedRepo(CONNECTION)

	//When
	err := repo.Insert(testFeed)

	//Then
	c.Assert(err, IsNil)
}

func (s *FeedRepoSuite) TestFeedRepo_All(c *C) {
	//Given
	repo := GetFeedRepo(CONNECTION)
	err := repo.Insert(testFeed)

	//When
	result, err := repo.All()

	//Then
	c.Assert(err, IsNil)
	c.Assert(len(result), Equals, 1)
}

func (s *FeedRepoSuite) TestFeedRepo_Update(c *C) {
	//Given
	repo := GetFeedRepo(CONNECTION)
	err := repo.Insert(testFeed)
	feeds, err := repo.All()
	feed := feeds[0]

	//When

	feed.Url = "http://localhost:54321/rss.xml"
	err = repo.Update(feed)
	result, _ := repo.All()

	//Then
	c.Assert(err, IsNil)
	c.Assert(len(result), Equals, 1)
	c.Assert(result[0].Url, Equals, "http://localhost:54321/rss.xml")
}

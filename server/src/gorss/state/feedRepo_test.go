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

var testFeeds = []domain.Feed{
	domain.Feed{
		Url: "http://localhost:12345/rss.xml", Tags: []string{"News", "Technology"}},
	domain.Feed{
		Url: "http://localhost:54321/rss.xml", Tags: []string{"News", "Technology"}}}

func (s *FeedRepoSuite) TestFeedRepo_Insert(c *C) {
	//Given
	repo := GetFeedRepo(CONNECTION)

	//When
	err := repo.Insert(testFeeds)

	//Then
	c.Assert(err, IsNil)
}

func (s *FeedRepoSuite) TestFeedRepo_All(c *C) {
	//Given
	repo := GetFeedRepo(CONNECTION)
	err := repo.Insert(testFeeds)

	//When
	result, err := repo.All()

	//Then
	c.Assert(err, IsNil)
	c.Assert(len(result), Equals, 2)
}

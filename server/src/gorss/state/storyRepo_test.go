package state

import (
	. "gopkg.in/check.v1"
	"gorss/domain"
	"time"
)

type StoryRepoSuite struct{}

var _ = Suite(&StoryRepoSuite{})

func (s *StoryRepoSuite) SetUpTest(c *C) {
	clearCollection(COLLECTION_STORIES)
}

var expectedDate = time.Now()

var stories = []domain.Story{
	domain.Story{Title: "A story", Id: "a_story", Date: expectedDate},
	domain.Story{Title: "Another story", Id: "another_story", Date: expectedDate}}

func (s *StoryRepoSuite) TestStoryRepo_Insert(c *C) {
	//Given
	repo := GetStoryRepo(CONNECTION)

	//When
	err := repo.Insert(stories)

	//Then
	c.Assert(err, IsNil)
}

func (s *StoryRepoSuite) TestStoryRepo_All(c *C) {
	//Given
	repo := GetStoryRepo(CONNECTION)
	err := repo.Insert(stories)

	//When
	result, err := repo.All()

	//Then
	c.Assert(err, IsNil)
	c.Assert(result, HasLen, 2)
}

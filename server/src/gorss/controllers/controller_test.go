package controllers

import (
	. "gopkg.in/check.v1"
	"gorss/state"
	"testing"
)

var CONNECTION = "localhost:27000"

func Test(t *testing.T) { TestingT(t) }

type ControllerSuite struct{}

var _ = Suite(&ControllerSuite{})

func (s *ControllerSuite) SetUpTest(c *C) {
	storyRepo = state.GetStoryRepo(CONNECTION)
	storyRepo.Clear()
}

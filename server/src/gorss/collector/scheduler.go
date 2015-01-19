package collector

import (
	"gorss/domain"
	"gorss/state"
	"log"
	. "time"
)

func ScheduleStoryCollection(feedStore state.FeedRepo, storyStore state.StoryRepo, sleepDuration Duration) {
	for {
		storyCollection(feedStore, storyStore)
		Sleep(sleepDuration)
	}
}

func ProcessFeed(feed domain.Feed, storyStore state.StoryRepo) {
	stories, err := LoadUrl(feed.Url, feed.Tags)
	if err != nil {
		log.Printf("Error Loading url: %v", err)
	} else {
		storyStore.Insert(stories)
	}
}

func storyCollection(feedStore state.FeedRepo, storyStore state.StoryRepo) {
	feeds, err := feedStore.All()

	if err != nil {
		log.Printf("Error loading feed list: %v", err)
	}
	if len(feeds) == 0 {
		log.Printf("No feeds to process")
	}
	for _, feed := range feeds {
		ProcessFeed(feed, storyStore)

	}
}

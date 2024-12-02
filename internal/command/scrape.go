package command

import (
	"context"
	"fmt"
	"log"

	"github.com/JP-Go/gator/internal/database"
)

func ScrapeFeeds(s *State) {
	ctx := context.Background()
	dbFeed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("Could not get next feed to fetch: %v", err)
		return
	}
	scrapeFeed(ctx, s.Db, dbFeed)
}

func scrapeFeed(ctx context.Context, db *database.Queries, f database.Feed) {
	err := db.MarkFeedFetched(ctx, f.ID)
	if err != nil {
		log.Printf("Could not mark feed %s as fetched: %v", f.Name, err)
		return
	}
	feed, err := FetchFeed(ctx, f.Url)
	if err != nil {
		log.Printf("Could not fetch feed %s: %v", f.Name, err)
		return
	}

	fmt.Printf("Feed %s has %d articles: \n", feed.Channel.Title, len(feed.Channel.Item))
	for _, item := range feed.Channel.Item {
		fmt.Println("  - " + item.Title)
	}

}

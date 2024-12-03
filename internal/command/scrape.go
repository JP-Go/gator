package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		wasSaved := savePost(ctx, db, item, f.ID)
		if !wasSaved {
			continue
		}
	}
}

func savePost(ctx context.Context, db *database.Queries, item RSSItem, feedID uuid.UUID) bool {
	fmt.Printf(" - Saving article %s to the database\n", item.Title)
	pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		log.Printf("   Could not save due to invalid date format %s\n", item.PubDate)
		return false
	}
	post, err := db.CreatePost(ctx, database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       item.Title,
		Url:         item.Link,
		FeedID:      feedID,
		Description: item.Description,
		PublishedAt: pubDate,
	})
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Constraint != "" {
			log.Printf(" - Article %s already in database\n", item.Title)
			return true
		} else {
			log.Printf(" - Error saving %s already in database: %s\n", item.Title, err)
			return false
		}
	}
	log.Printf(" | Saved post %s\n", post.Title)
	return true
}

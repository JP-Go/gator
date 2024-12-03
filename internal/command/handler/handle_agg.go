package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/database"
)

func handleAgg(s *command.State, cmd command.Command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected 1, got %d", len(cmd.Arguments))
	}
	duration, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", duration)
	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		command.ScrapeFeeds(s)
	}
}

func handleBrowse(ctx context.Context, s *command.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("Invalid number of arguments. Expected 1, got %d", len(cmd.Arguments))
	}
	limit := 2
	if len(cmd.Arguments) == 1 {
		newLimit, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			fmt.Printf("Invalid limit %s, defaulting to two posts", cmd.Arguments[0])
		}
		limit = newLimit
	}

	posts, err := s.Db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Could not get posts due to: %w", err)
	}
	if len(posts) == 0 {
		fmt.Println("Following no feeds")
		return nil
	}
	fmt.Println("Recent posts")
	for i := 0; i < limit; i++ {
		post := posts[i]
		fmt.Println("==============================")
		fmt.Printf(" %s posted %s ago \n", post.FeedName, time.Since(post.PublishedAt).Truncate(time.Second))
		fmt.Printf(" - %s at %s \n", post.Title, post.Url)
		charLimit := 64
		if len(post.Description) < charLimit {
			charLimit = len(post.Description) - 1
		}
		fmt.Printf(" - %s \n", post.Description[0:charLimit])
		fmt.Println("==============================")
	}
	return nil
}

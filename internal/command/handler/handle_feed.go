package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(ctx context.Context, s *command.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("Expecting 2 arguments <title> <url>, got %v", len(cmd.Arguments))
	}
	name, url := cmd.Arguments[0], cmd.Arguments[1]
	feed, err := s.Db.AddFeed(ctx, database.AddFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed created: %s at %s", feed.Name, feed.Url)

	return nil
}

func handleListFeeds(s *command.State, cmd command.Command) error {
	if len(cmd.Arguments) != 0 {
		return errors.New("Unexpected arguments. Expected none.")
	}

	feedsWithUserName, err := s.Db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feeds:")
	for _, feedWithUserName := range feedsWithUserName {
		fmt.Printf("- Name: %s\n  URL: %s\n  Added By: %s\n",
			feedWithUserName.Name,
			feedWithUserName.Url,
			feedWithUserName.UserName,
		)
	}

	return nil
}

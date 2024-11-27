package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/JP-Go/gator/internal/command"
	internal_errors "github.com/JP-Go/gator/internal/command/errors"
	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(ctx context.Context, s *command.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Expecting 1 arguments <url>, got %v", len(cmd.Arguments))
	}
	url := cmd.Arguments[0]

	feed, err := s.Db.FindFeedByURL(ctx, url)
	if errors.Is(err, sql.ErrNoRows) {
		return internal_errors.ErrFeedNotFound
	}
	if err != nil {
		return err
	}

	follow, err := s.Db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("User %s following feed %s. Good Reading\n", follow.UserName, follow.FeedName)

	return nil
}

func handleFollowing(ctx context.Context, s *command.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("Expecting no arguments, got %v", len(cmd.Arguments))
	}

	feedFollows, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User %v is following these feeds: \n", user.Name)
	if len(feedFollows) == 0 {
		fmt.Println("  No feeds")
		return nil
	} else {
		for _, feedFollow := range feedFollows {
			fmt.Printf(" - %v at %v\n", feedFollow.FeedName, feedFollow.FeedUrl)
		}
	}

	return nil
}

func handleUnfollow(ctx context.Context, s *command.State, cmd command.Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Expecting two arguments <url>, got %v", len(cmd.Arguments))
	}
	url := cmd.Arguments[0]
	feed, err := s.Db.FindFeedByURL(ctx, url)
	if errors.Is(err, sql.ErrNoRows) {
		return internal_errors.ErrFeedNotFound
	}
	if err != nil {
		return err
	}
	err = s.Db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Could not unfollow feed due to %w", err)
	}
	fmt.Printf("%s unfollowed successfully", feed.Name)
	return nil
}

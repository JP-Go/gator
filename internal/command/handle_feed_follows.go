package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"time"

	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command) error {

	if s.config.CurrentUserName == "" {
		return fmt.Errorf("Not logged in. Cannot follow feed")
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Expecting 1 arguments <url>, got %v", len(cmd.arguments))
	}
	url := cmd.arguments[0]
	ctx := context.Background()

	var feedToFollow *database.Feed
	var userFollowing *database.User
	var fetchErr error
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		feed, err := s.db.FindFeedByURL(ctx, url)
		if errors.Is(err, sql.ErrNoRows) {
			fetchErr = errors.New("Feed not found")
			return
		}
		if err != nil {
			fetchErr = err
			return
		}
		feedToFollow = &feed
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		user, err := s.db.FindUserByName(ctx, s.config.CurrentUserName)
		if errors.Is(err, sql.ErrNoRows) {
			fetchErr = errors.New("User not found")
			return
		}
		if err != nil {
			fetchErr = err
			return
		}
		userFollowing = &user
	}()

	wg.Wait()
	if fetchErr != nil {
		return fetchErr
	}

	follow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feedToFollow.ID,
		UserID:    userFollowing.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("User %s following feed %s. Good Reading\n", follow.UserName, follow.FeedName)

	return nil
}

func handleFollowing(s *state, cmd command) error {
	if s.config.CurrentUserName == "" {
		return fmt.Errorf("Not logged in. Cannot see following")
	}

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Expecting no arguments, got %v", len(cmd.arguments))
	}

	ctx := context.Background()
	user, err := s.db.FindUserByName(ctx, s.config.CurrentUserName)
	if err != nil {
		return err
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
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

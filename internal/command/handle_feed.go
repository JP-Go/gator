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

func handleListFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("Unexpected arguments. Expected none.")
	}

	feedsWithUserName, err := s.db.GetFeedsWithUserName(context.Background())
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

func handleAddFeed(s *state, cmd command) error {
	ctx := context.Background()

	if s.config.CurrentUserName == "" {
		return fmt.Errorf("Not logged in. Cannot add feed")
	}

	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Expecting 2 arguments <title> <url>, got %v", len(cmd.arguments))
	}
	name, url := cmd.arguments[0], cmd.arguments[1]

	user, err := s.db.FindUserByName(ctx, s.config.CurrentUserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Could not find current user in the database. Exiting")
		}
		return err
	}
	feed, err := s.db.AddFeed(ctx, database.AddFeedParams{
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
	fmt.Printf("Feed created: %s at %s", feed.Name, feed.Url)

	return nil
}

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

package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
)

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

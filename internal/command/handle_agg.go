package command

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Could not fetch feed: %w", err)
	}
	fmt.Println("Feed:")
	fmt.Printf(" %+v\n", feed)

	return nil
}

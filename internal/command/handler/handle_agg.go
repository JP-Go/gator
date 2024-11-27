package handler

import (
	"context"
	"fmt"

	"github.com/JP-Go/gator/internal/command"
)

func handleAgg(s *command.State, cmd command.Command) error {

	feed, err := command.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Could not fetch feed: %w", err)
	}
	fmt.Println("Feed:")
	fmt.Printf(" %+v\n", feed)

	return nil
}

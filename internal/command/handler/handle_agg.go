package handler

import (
	"fmt"
	"time"

	"github.com/JP-Go/gator/internal/command"
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
		command.ScrapeFeed(s)
	}
}

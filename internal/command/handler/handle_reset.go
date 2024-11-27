package handler

import (
	"context"
	"fmt"

	"github.com/JP-Go/gator/internal/command"
)

func handleReset(s *command.State, cmd command.Command) error {
	err := s.Db.DeleteAllUsers(context.Background())

	if cmd.Name != "reset" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.Name))
	}

	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("Invalid number of arguments. Expected no arguments, given %d",
			len(cmd.Arguments),
		)
	}
	if err != nil {
		return fmt.Errorf("Could not delete all users %w", err)
	}
	fmt.Println("Users deleted")
	return nil
}

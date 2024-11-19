package command

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())

	if cmd.name != "reset" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.name))
	}

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Invalid number of arguments. Expected no arguments, given %d",
			len(cmd.arguments),
		)
	}
	if err != nil {
		return fmt.Errorf("Could not delete all users %w", err)
	}
	fmt.Println("Users deleted")
	return nil
}

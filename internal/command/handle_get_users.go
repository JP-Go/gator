package command

import (
	"context"
	"fmt"
	"os"

	"github.com/lib/pq"
)

func handleGetUsers(s *state, cmd command) error {

	if cmd.name != "users" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.name))
	}

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("Invalid number of arguments. Expected no arguments, given %d",
			len(cmd.arguments),
		)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		if err.(*pq.Error).Constraint == "users_name_key" {
			fmt.Println("Error: username already in use. Choose another one.")
			os.Exit(1)
		}
	}

	for _, user := range users {
		fmt.Printf("* %v ", user.Name)
		if user.Name == s.config.CurrentUserName {
			fmt.Println("(current)")
		} else {
			fmt.Println()
		}
	}
	return nil
}

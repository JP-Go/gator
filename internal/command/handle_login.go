package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {

	if cmd.name != "login" {
		panic(fmt.Errorf("Handle login called for %v command. Invalid state. Exiting", cmd.name))
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected one argument [username], given %d",
			len(cmd.arguments),
		)
	}

	name := cmd.arguments[0]
	user, err := s.db.FindUserByName(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("User not found.")
		}
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Could not login due to: %v", err)
	}
	fmt.Println("User logged in: " + s.config.CurrentUserName)
	return nil
}

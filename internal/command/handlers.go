package command

import (
	"fmt"
)

func HandleLogin(s *state, cmd command) error {

	if cmd.name != "login" {
		return fmt.Errorf("Invalid command handler given handleLogin. given: %s", cmd.name)
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected one argument [username], given %d",
			len(cmd.arguments),
		)
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Could not login due to: %v", err)
	}
	fmt.Println("User logged in: " + s.config.CurrentUserName)
	return nil
}

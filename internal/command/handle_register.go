package command

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handleRegister(s *state, cmd command) error {

	if cmd.name != "register" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.name))
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected one argument [new username], given %d",
			len(cmd.arguments),
		)
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		if err.(*pq.Error).Constraint == "users_name_key" {
			fmt.Println("Error: username already in use. Choose another one.")
			os.Exit(1)
		}
	}

	s.config.SetUser(user.Name)
	fmt.Println("User registered and logged in: " + s.config.CurrentUserName)
	return nil
}

package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"

	"database/sql"

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

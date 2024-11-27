package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"database/sql"
)

func handleGetUsers(s *command.State, cmd command.Command) error {

	if cmd.Name != "users" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.Name))
	}

	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("Invalid number of arguments. Expected no arguments, given %d",
			len(cmd.Arguments),
		)
	}
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		if err.(*pq.Error).Constraint == "users_name_key" {
			fmt.Println("Error: username already in use. Choose another one.")
			os.Exit(1)
		}
	}

	for _, user := range users {
		fmt.Printf("* %v ", user.Name)
		if user.Name == s.Config.CurrentUserName {
			fmt.Println("(current)")
		} else {
			fmt.Println()
		}
	}
	return nil
}

func handleLogin(s *command.State, cmd command.Command) error {

	if cmd.Name != "login" {
		panic(fmt.Errorf("Handle login called for %v command. Invalid state. Exiting", cmd.Name))
	}

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected one argument [username], given %d",
			len(cmd.Arguments),
		)
	}

	name := cmd.Arguments[0]
	user, err := s.Db.FindUserByName(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("User not found.")
		}
	}
	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Could not login due to: %v", err)
	}
	fmt.Println("User logged in: " + s.Config.CurrentUserName)
	return nil

}

func handleRegister(s *command.State, cmd command.Command) error {

	if cmd.Name != "register" {
		panic(fmt.Errorf("Handle register called for %v command. Invalid state. Exiting", cmd.Name))
	}

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("Invalid number of arguments. Expected one argument [new username], given %d",
			len(cmd.Arguments),
		)
	}
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	})
	if err != nil {
		if err.(*pq.Error).Constraint == "users_name_key" {
			fmt.Println("Error: username already in use. Choose another one.")
			os.Exit(1)
		}
	}

	s.Config.SetUser(user.Name)
	fmt.Println("User registered and logged in: " + s.Config.CurrentUserName)
	return nil
}

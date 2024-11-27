package middleware

import (
	"context"
	"database/sql"
	"errors"

	"github.com/JP-Go/gator/internal/command"
	internal_errors "github.com/JP-Go/gator/internal/command/errors"
	"github.com/JP-Go/gator/internal/database"
)

type HandlerWithUser func(context.Context, *command.State, command.Command, database.User) error

func MiddlewareLoggedIn(handler HandlerWithUser) command.CommandHandler {
	return func(s *command.State, cmd command.Command) error {
		if s.Config.CurrentUserName == "" {
			return internal_errors.ErrNotLoggedIn
		}
		ctx := context.Background()

		user, err := s.Db.FindUserByName(ctx, s.Config.CurrentUserName)
		if errors.Is(err, sql.ErrNoRows) {
			return internal_errors.ErrFeedNotFound
		}
		if err != nil {
			return err
		}

		return handler(ctx, s, cmd, user)
	}
}

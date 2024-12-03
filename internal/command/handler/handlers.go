package handler

import (
	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/command/middleware"
)

func RegisterCommands() *command.Commands {
	cmds := command.NewCommands()
	cmds.Register("login", handleLogin)
	cmds.Register("register", handleRegister)
	cmds.Register("reset", handleReset)
	cmds.Register("users", handleGetUsers)
	cmds.Register("agg", handleAgg)
	cmds.Register("addfeed", middleware.MiddlewareLoggedIn(handleAddFeed))
	cmds.Register("feeds", handleListFeeds)
	cmds.Register("follow", middleware.MiddlewareLoggedIn(handleFollow))
	cmds.Register("following", middleware.MiddlewareLoggedIn(handleFollowing))
	cmds.Register("unfollow", middleware.MiddlewareLoggedIn(handleUnfollow))
	cmds.Register("browse", middleware.MiddlewareLoggedIn(handleBrowse))
	return cmds
}

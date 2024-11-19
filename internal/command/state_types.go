package command

import (
	"github.com/JP-Go/gator/internal/config"
	"github.com/JP-Go/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

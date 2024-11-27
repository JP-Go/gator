package command

import (
	"github.com/JP-Go/gator/internal/config"
	"github.com/JP-Go/gator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}

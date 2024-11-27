package command

import (
	"github.com/JP-Go/gator/internal/config"
	"github.com/JP-Go/gator/internal/database"
)

func NewState(config *config.Config, queries *database.Queries) *State {
	return &State{
		Config: config,
		Db:     queries,
	}
}

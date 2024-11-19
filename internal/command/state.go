package command

import (
	"github.com/JP-Go/gator/internal/config"
	"github.com/JP-Go/gator/internal/database"
)

func NewState(config *config.Config, queries *database.Queries) *state {
	return &state{
		config: config,
		db:     queries,
	}
}

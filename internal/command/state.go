package command

import "github.com/JP-Go/gator/internal/config"

func NewState(config *config.Config) *state {
	return &state{
		config: config,
	}
}

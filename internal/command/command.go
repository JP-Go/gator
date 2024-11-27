package command

import (
	"errors"
)

func newCommand(name string, arguments []string) Command {
	return Command{
		Name:      name,
		Arguments: arguments,
	}
}

func NewCommand(name string, arguments []string) Command {
	return newCommand(name, arguments)
}

func NewCommands() *Commands {
	return &Commands{
		handlers: map[string]CommandHandler{},
	}
}

func (c *Commands) Register(name string, f CommandHandler) {
	c.handlers[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}
	return handler(s, cmd)
}

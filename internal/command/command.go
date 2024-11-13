package command

import "errors"

func newCommand(name string, arguments []string) command {
	return command{
		name:      name,
		arguments: arguments,
	}
}

func NewCommand(name string, arguments []string) command {
	return command{name: name, arguments: arguments}
}

func NewCommands() *commands {
	return &commands{
		handlers: map[string]commandHandler{},
	}
}

func (c *commands) Register(name string, f commandHandler) {
	c.handlers[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}
	return handler(s, cmd)
}

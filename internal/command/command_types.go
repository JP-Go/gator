package command

type command struct {
	name      string
	arguments []string
}
type commandHandler = func(*state, command) error

type commands struct {
	handlers map[string]commandHandler
}

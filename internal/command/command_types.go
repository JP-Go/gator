package command

type Command struct {
	Name      string
	Arguments []string
}
type CommandHandler = func(*State, Command) error

type Commands struct {
	handlers map[string]CommandHandler
}

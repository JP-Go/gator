package command

func RegisterCommands() *commands {
	cmds := NewCommands()
	cmds.Register("login", handleLogin)
	cmds.Register("register", handleRegister)
	cmds.Register("reset", handleReset)
	cmds.Register("users", handleGetUsers)
	cmds.Register("agg", handleAgg)
	return cmds
}

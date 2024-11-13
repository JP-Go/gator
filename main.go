package main

import (
	"fmt"
	"os"

	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	state := command.NewState(&cfg)
	if state == nil {
		return
	}
	cmds := command.NewCommands()
	cmds.Register("login", command.HandleLogin)
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	if err := cmds.Run(state, command.NewCommand(args[1], args[2:])); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

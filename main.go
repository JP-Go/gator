package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JP-Go/gator/internal/command"
	"github.com/JP-Go/gator/internal/command/handler"
	"github.com/JP-Go/gator/internal/config"
	"github.com/JP-Go/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)
	state := command.NewState(cfg, dbQueries)
	if state == nil {
		return
	}
	cmds := handler.RegisterCommands()
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	if err := cmds.Run(state, command.NewCommand(args[1], args[2:])); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

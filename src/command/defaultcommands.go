package command

import (
	"fmt"
	"os"

	"github.com/qdiaps/tirf-assistant/src/types"
)

var (
	Commands map[string]types.Command
)

func InitDefaultCommands() {
	Commands = map[string]types.Command{
		"help":    {help, "Show all commands."},
		"version": {version, "Show version."},
		"exit":    {exit, "Exit the application."},
	}
}

func help() {
	fmt.Println("List all commands:")
	for name, cmd := range Commands {
		fmt.Printf("  %-40s%s\n\n", name, cmd.Desc)
	}
}

func version() {
	fmt.Printf("Version: v%s\n", types.GetVersion())
}

func exit() {
	os.Exit(0)
}

package src

import (
	"fmt"

	"github.com/qdiaps/tirf-assistant/src/command"
	"github.com/qdiaps/tirf-assistant/src/types"
)

func Run() {
	config := LoadConfig()
	fmt.Printf("%s v%s\n", config.Name, types.GetVersion())

	command.InitDefaultCommands()
	command.HandleUserInput()
}

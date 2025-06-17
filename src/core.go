package src

import (
	"fmt"

	"github.com/qdiaps/tirf-assistant/src/command"
	"github.com/qdiaps/tirf-assistant/src/config"
	"github.com/qdiaps/tirf-assistant/src/module"
	"github.com/qdiaps/tirf-assistant/src/types"
)

func Run() {
	cfg := config.GetConfig()
	fmt.Printf("%s v%s\n", cfg.Name, types.GetVersion())

	module.LoadAllModules()

	command.InitDefaultCommands()
	command.HandleUserInput()
}

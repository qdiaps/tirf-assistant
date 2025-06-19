package module

import (
	"github.com/qdiaps/tirf-assistant/src/types"
)

var (
	cfg *types.Config

	startedModules = make(map[string]types.Module)
	loadedModules  = make(map[string]types.Module)
)

package module

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/qdiaps/tirf-assistant/src/config"
	"github.com/qdiaps/tirf-assistant/src/types"
	"github.com/qdiaps/tirf-assistant/src/utils"
)

var (
	cfg *types.Config

	startedModules = make(map[string]types.Module)
	loadedModules  = make(map[string]types.Module)
)

func LoadAllModules() {
	if cfg == nil {
		cfg = config.GetConfig()
	}

	if len(cfg.Modules) != 0 {
		loadModulesFromConfig()
	}
	loadModulesFromDir()

	fmt.Println("Modules loaded:")
	for name := range loadedModules {
		fmt.Printf("  %s\n", name)
	}
}

func LoadModule(name string, params *types.Module) {
	if _, exists := loadedModules[name]; exists == false {
		loadedModules[name] = *params
	} else {
		// TODO: log a warning that the module has already been loaded
	}
}

func loadModulesFromConfig() {
	for name, module := range cfg.Modules {
		LoadModule(name, &module)
	}
}

func loadModulesFromDir() {
	files, err := os.ReadDir(types.PathToModules)
	if err != nil {
		log.Fatalf("Error reading directory: %s", err)
	}

	for _, file := range files {
		if file.IsDir() == false {
			// TODO: log a warning that there is an extra file in the modules
			continue
		}

		manifest := types.PathToModules + file.Name() + "/" + types.ModuleManifestName
		if utils.IsExist(manifest) == false {
			// TODO: log a warning that there is no `module.manifest.json` in the module
			continue
		}

		module := loadManifest(manifest)
		validateModuleData(module, file.Name())
		saveNewModule(module)
		LoadModule(module.Name, module)
	}
}

func loadManifest(path string) *types.Module {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Read file error: %s", err)
	}

	m := &types.Module{}
	if err = json.Unmarshal(data, m); err != nil {
		log.Fatalf("Encode config file error: %s", err)
	}

	return m
}

func validateModuleData(module *types.Module, name string) {
	module.Path = types.PathToModules + name + "/"
	module.IsActive = true
}

func saveNewModule(module *types.Module) {
	if _, exists := cfg.Modules[module.Name]; exists == false {
		cfg.Modules[module.Name] = *module
		config.SaveConfig()
	}
}

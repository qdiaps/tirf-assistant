package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/qdiaps/tirf-assistant/src/types"
	"github.com/qdiaps/tirf-assistant/src/utils"
)

var (
	config *types.Config
)

func GetConfig() *types.Config {
	if config == nil {
		return loadConfig()
	}

	return config
}

func SaveConfig() {
	jsonConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Decode config error: %s", err)
	}

	err = os.WriteFile(types.PathToConfig, jsonConfig, 0644)
	if err != nil {
		log.Fatalf("Write file error: %s", err)
	}
}

func loadConfig() *types.Config {
	if utils.IsExist(types.PathToConfig) == false {
		return createDefaultConfig()
	}

	data, err := os.ReadFile(types.PathToConfig)
	if err != nil {
		log.Fatalf("Read file error: %s", err)
	}

	config = &types.Config{}
	if err = json.Unmarshal(data, config); err != nil {
		log.Fatalf("Encode config file error: %s", err)
	}

	return config
}

func createDefaultConfig() *types.Config {
	config = &types.Config{
		Name:    types.Name,
		Modules: make(map[string]types.Module),
	}

	SaveConfig()

	return config
}

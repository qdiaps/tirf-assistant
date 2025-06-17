package src

import (
	"encoding/json"
	"log"
	"os"

	"github.com/qdiaps/tirf-assistant/src/types"
)

func LoadConfig() *types.Config {
	if isExist(types.PathToConfig) == false {
		return createDefaultConfig()
	}

	data, err := os.ReadFile(types.PathToConfig)
	if err != nil {
		log.Fatalf("Read file error: %s", err)
	}

	var config types.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Encode config file error: %s", err)
	}

	return &config
}

func isExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func createDefaultConfig() *types.Config {
	config := &types.Config{
		Name: types.DefaultName,
	}

	jsonConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Decode config error: %s", err)
	}

	err = os.WriteFile(types.PathToConfig, jsonConfig, 0644)
	if err != nil {
		log.Fatalf("Write file error: %s", err)
	}

	return config
}

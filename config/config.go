package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadConfigFile Function to read and parse the config file
func ReadConfigFile(configFilePath string) (map[string]interface{}, error) {
	configFileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var configData map[string]interface{}
	err = json.Unmarshal(configFileContent, &configData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return configData, nil
}

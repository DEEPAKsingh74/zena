package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type ProviderConfig struct {
	Key     string `json:"key"`
	Default bool   `json:"default"`
}

type Config struct {
	OpenAI    ProviderConfig `json:"openai"`
	Anthropic ProviderConfig `json:"anthropic"`
	Gemini    ProviderConfig `json:"gemini"`
}

/*
	getConfigPath returns the path to the configuration file based on the operating system.
	It uses the user's home directory and constructs a path that is appropriate for the OS.
	On Windows, it returns a path in the AppData directory.
	On Unix-like systems, it returns a path in the user's .config directory.
*/

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	if runtime.GOOS == "windows" {
		return filepath.Join(homeDir, "AppData", "Local", "zena", "config.json")
	}
	return filepath.Join(homeDir, ".config", "zena", "config.json")
}

/*
	LoadConfig reads the configuration file from the user's home directory.
	It decodes the JSON content into a Config struct and returns it.
	If the file does not exist or there is an error reading it, it returns an empty Config
	struct and the error.
	It uses the getConfigPath function to determine the correct path to the configuration file.
	If the file is not found, it returns an empty Config struct and no error.
	It uses the json package to decode the JSON content into the Config struct.

*/

func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()
	file, err := os.Open(configPath)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}


func SaveConfig(config *Config) error {
	configPath := GetConfigPath()
	dir := filepath.Dir(configPath)

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}
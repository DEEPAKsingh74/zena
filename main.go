package main

import (
	"os"

	"github.com/DEEPAKsingh74/zena/cmd"
	"github.com/DEEPAKsingh74/zena/internal/config"
)

// const Version = "0.1.0"

func main() {
	_, err := config.LoadConfig()

	if err != nil {
		if os.IsNotExist(err) {
			// create a default configuration file
			defaultConfig := &config.Config{
				OpenAI: config.ProviderConfig{
					Key:     "",
					Default: true,
				},
				Anthropic: config.ProviderConfig{
					Key:     "",
					Default: false,
				},
				Gemini: config.ProviderConfig{
					Key:     "",
					Default: false,
				},
			}
			err = config.SaveConfig(defaultConfig)
			if err != nil {
				panic("❌ Failed to create default configuration: " + err.Error())
			}

			println("✅ Default configuration created at:", config.GetConfigPath())
		} else {
			panic("❌ Failed to load configuration: " + err.Error())
		}
	}

	cmd.Execute()
}
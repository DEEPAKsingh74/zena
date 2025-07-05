package helpers

import (
	"errors"

	"github.com/DEEPAKsingh74/zena/internal/config"
)


func MarkDefault(provider string) (string, error) {
	if provider == "" {
		return "", nil 
	}

	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	// Reset all providers to non-default
	cfg.OpenAI.Default = false
	cfg.Anthropic.Default = false
	cfg.Gemini.Default = false

	// Mark the specified provider as default
	switch provider {
	case "openai":
		cfg.OpenAI.Default = true
	case "anthropic":
		cfg.Anthropic.Default = true
	case "gemini":
		cfg.Gemini.Default = true
	default:
		return "", errors.New("unsupported provider: " + provider)
	}

	// Save the updated configuration
	if err := config.SaveConfig(cfg); err != nil {
		return "", err
	}

	return provider, nil
}
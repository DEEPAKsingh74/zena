package helpers

import (
	"errors"
	"github.com/DEEPAKsingh74/zena/internal/config"
)

func GetAiApiKey() (string, string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", "", err
	}

	// Check for default provider in order
	if cfg.OpenAI.Default {
		return "openai", cfg.OpenAI.Key, nil
	}
	if cfg.Anthropic.Default {
		return "anthropic", cfg.Anthropic.Key, nil
	}
	if cfg.Gemini.Default {
		return "gemini", cfg.Gemini.Key, nil
	}

	return "", "", errors.New("no default provider is set")
}
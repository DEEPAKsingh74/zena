package validators

import (
	"errors"
	"github.com/DEEPAKsingh74/zena/internal/config"
)

// ValidateKeyProvider checks if the provided key and provider are valid.

func ValidateKeyProvider(provider, key string) error {
	if provider == "" {
		return errors.New("provider cannot be empty")
	}

	if key == "" {
		return errors.New("API key cannot be empty")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	switch provider {
	case "openai":
		if cfg.OpenAI.Key != key {
			return errors.New("invalid OpenAI API key")
		}
	case "anthropic":
		if cfg.Anthropic.Key != key {
			return errors.New("invalid Anthropic API key")
		}
	case "gemini":
		if cfg.Gemini.Key != key {
			return errors.New("invalid Gemini API key")
		}
	default:
		return errors.New("unsupported provider: " + provider)
	}

	return nil
}
package ai

import (
	"errors"

)

func FetchResponse(query string, provider string, apiKey string) ([]*AIResponse, error) {

	// Get the response from the AI provider

	switch provider {
	case "openai":
		return FetchResponseOpenAI(query, apiKey)
	case "anthropic":
		// return FetchResponseAnthropic(query, apiKey)
	case "gemini":
		return FetchResponseGemini(query, apiKey)
	default:
		return nil, errors.New("unsupported AI provider: " + provider)
	}

	return nil, errors.New("unsupported AI provider: " + provider)
}
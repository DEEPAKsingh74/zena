package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DEEPAKsingh74/zena/internal/utils/constants"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"

type openAIRequest struct {
	Model    string          `json:"model"`
	Messages []openAIMessage `json:"messages"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
}

// FetchResponseOpenAI sends the query to OpenAI and returns the response.
func FetchResponseOpenAI(query string, apiKey string) ([]*AIResponse, error) {
	payload := openAIRequest{
		Model: "gpt-3.5-turbo", // or "gpt-4" if needed
		Messages: []openAIMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf(constants.OpenAIPromptTemplate, query),
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}


	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(resp.Body)
		return nil, errors.New("OpenAI API error: " + string(resBody))
	}

	var apiResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if len(apiResp.Choices) == 0 {
		return nil, errors.New("no choices returned from OpenAI")
	}

	message := apiResp.Choices[0].Message.Content

	return []*AIResponse{
		{
			Text:  message,
			Color: "blue",
			Type:  "Note",
		},
	}, nil
}

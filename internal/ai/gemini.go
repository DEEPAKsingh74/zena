package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/DEEPAKsingh74/zena/internal/utils/constants"
)

const geminiURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key="

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Role  string         `json:"role"`
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// FetchResponseGemini sends a prompt to Gemini and returns AIResponse chunks.
func FetchResponseGemini(query string, apiKey string) ([]*AIResponse, error) {
	url := geminiURL + apiKey

	payload := geminiRequest{
		Contents: []geminiContent{
			{
				Role: "user",
				Parts: []geminiPart{
					{Text: fmt.Sprintf(constants.OpenAIPromptTemplate, query)},
				},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// fmt.Println("🔍 Sending request to Gemini...", string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(resp.Body)
		return nil, errors.New("Gemini API error: " + string(resBody))
	}

	var gemResp geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&gemResp); err != nil {
		return nil, err
	}

	if len(gemResp.Candidates) == 0 || len(gemResp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("no response returned from Gemini")
	}

	text := gemResp.Candidates[0].Content.Parts[0].Text

	// Strip markdown code block fences if present
	if len(text) > 0 {
		text = strings.TrimSpace(text)
		if strings.HasPrefix(text, "```") {
			text = strings.TrimPrefix(text, "```json")
			text = strings.TrimPrefix(text, "```") // fallback if not labeled json
			text = strings.TrimSuffix(text, "```")
			text = strings.TrimSpace(text)
		}
	}

	// In your case, the response is a JSON array string, so unmarshal it.
	var parsed []*AIResponse
	if err := json.Unmarshal([]byte(text), &parsed); err != nil {
		return nil, fmt.Errorf("error parsing Gemini response JSON: %w\nRaw response: %s", err, text)
	}

	return parsed, nil
}

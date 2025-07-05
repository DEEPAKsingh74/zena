package ai


type AIResponse struct {
	Text          string `json:"text"`
	Color         string `json:"color,omitempty"`
	Type          string `json:"type"` // Note, Warning, Error, Command
	RevertCommand string `json:"revertCommand,omitempty"`
}


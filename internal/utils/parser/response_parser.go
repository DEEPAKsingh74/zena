package parser

import (
	"fmt"
	"strings"

	"github.com/DEEPAKsingh74/zena/internal/ai"
	"github.com/DEEPAKsingh74/zena/internal/utils/constants"
)

func ParseAIResponse(responses []*ai.AIResponse) (string, error) {
	var sb strings.Builder

	for i, resp := range responses {
		// Fallback to default color
		color := resp.Color
		if color == "" {
			color = constants.ColorBlue
		}

		colorCode := constants.ColorMap[strings.ToLower(color)]
		reset := constants.ColorReset

		// Header
		if strings.ToLower(resp.Type) != "note" {
			sb.WriteString(fmt.Sprintf("%s[%d] [%s]%s\n", colorCode, i+1, resp.Type, reset))
		}

		// Main Text
		sb.WriteString(fmt.Sprintf("%s%s%s\n", colorCode, resp.Text, reset))

		// Optional RevertCommand
		if resp.RevertCommand != "" {
			sb.WriteString(fmt.Sprintf("%s↩ Revert: %s%s\n", constants.ColorYellow, resp.RevertCommand, reset))
		}

		sb.WriteString("\n")
	}

	return sb.String(), nil
}

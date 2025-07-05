package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DEEPAKsingh74/zena/internal/ai"
	"github.com/DEEPAKsingh74/zena/internal/utils/helpers"
	"github.com/DEEPAKsingh74/zena/internal/utils/parser"
	"github.com/DEEPAKsingh74/zena/internal/validators"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zena [query]",
	Short: "Zena is your AI-powered CLI assistant",
	Long: `Zena is an intelligent command-line assistant that uses AI models 
(OpenAI, Anthropic, Gemini) to help you generate, understand, and run commands 
or answer natural language queries directly from your terminal.

Examples:
  zena "how to create a zip file in linux"
  zena "give me a curl command to send a POST request with JSON"
  zena "remove a directory recursively in bash"

Zena will fetch AI-generated structured output and format it for your terminal, 
with support for command suggestions, warnings, notes, and revert instructions.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userQuery := strings.Join(args, " ")

		// 🔑 Get active API key and provider from config
		provider, apiKey, err := helpers.GetAiApiKey()
		if err != nil {
			fmt.Println("❌ Failed to retrieve AI API key:", err)
			return
		}

		// ✅ Validate the provider and its key
		if err := validators.ValidateKeyProvider(provider, apiKey); err != nil {
			fmt.Println("❌ Invalid API key or unsupported provider:", err)
			return
		}

		// 🤖 Fetch AI response
		response, err := ai.FetchResponse(userQuery, provider, apiKey)
		if err != nil {
			fmt.Println("❌ Error fetching AI response:", err)
			return
		}

		// 🎨 Parse and display the AI response
		output, err := parser.ParseAIResponse(response)
		if err != nil {
			fmt.Println("❌ Failed to parse AI response:", err)
			return
		}

		fmt.Println(output)
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// init is used for command setup and global flags
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
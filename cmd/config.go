package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DEEPAKsingh74/zena/internal/config"
	"github.com/spf13/cobra"
)

var listFlag bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Zena AI provider API keys and preferences",
	Long: `The 'config' command allows you to manage your API keys 
for AI providers (OpenAI, Anthropic, Gemini) and set your preferred default provider.

Examples:
  zena config set openai <your-api-key>     Set your OpenAI API key
  zena config set default openai            Set OpenAI as your default provider
  zena config --list                        List all configured providers and their default status`,
	Run: func(cmd *cobra.Command, args []string) {
		if listFlag {
			listConfig()
		} else {
			_ = cmd.Help()
		}
	},
}

var setCmd = &cobra.Command{
	Use:   "set [provider|default] [value]",
	Short: "Set API key or default provider",
	Long: `Set API keys or change the default provider used by Zena.

Usage:
  zena config set openai <your-api-key>       Sets the API key for OpenAI
  zena config set anthropic <your-api-key>    Sets the API key for Anthropic
  zena config set gemini <your-api-key>       Sets the API key for Gemini
  zena config set default openai              Sets OpenAI as the default provider

The default provider is used when no provider is explicitly specified.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		keyOrDefault := strings.ToLower(args[0])
		value := args[1]

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("❌ Failed to load config:", err)
			os.Exit(1)
		}

		switch keyOrDefault {
		case "openai", "anthropic", "gemini":
			setAPIKey(cfg, keyOrDefault, value)
		case "default":
			setDefaultProvider(cfg, strings.ToLower(value))
		default:
			fmt.Println("❌ Invalid usage. Try:")
			fmt.Println("  zena config set openai <apiKey>")
			fmt.Println("  zena config set default <provider>")
		}
	},
}

func setAPIKey(cfg *config.Config, provider, key string) {
	switch provider {
	case "openai":
		cfg.OpenAI.Key = key
	case "anthropic":
		cfg.Anthropic.Key = key
	case "gemini":
		cfg.Gemini.Key = key
	}

	if err := config.SaveConfig(cfg); err != nil {
		fmt.Println("❌ Failed to save config:", err)
		os.Exit(1)
	}
	fmt.Printf("✅ API key set for '%s'\n", provider)
}

func setDefaultProvider(cfg *config.Config, provider string) {
	// Reset all defaults
	cfg.OpenAI.Default = false
	cfg.Anthropic.Default = false
	cfg.Gemini.Default = false

	switch provider {
	case "openai":
		cfg.OpenAI.Default = true
	case "anthropic":
		cfg.Anthropic.Default = true
	case "gemini":
		cfg.Gemini.Default = true
	default:
		fmt.Println("❌ Invalid provider. Choose from: openai, anthropic, gemini")
		return
	}

	if err := config.SaveConfig(cfg); err != nil {
		fmt.Println("❌ Failed to save config:", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Default provider set to '%s'\n", provider)
}

func listConfig() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("❌ Failed to load config:", err)
		os.Exit(1)
	}

	fmt.Println("📦 Configured Providers:")
	if cfg.OpenAI.Key != "" {
		fmt.Printf("  - openai     ✅  (default: %v)\n", cfg.OpenAI.Default)
	}
	if cfg.Anthropic.Key != "" {
		fmt.Printf("  - anthropic  ✅  (default: %v)\n", cfg.Anthropic.Default)
	}
	if cfg.Gemini.Key != "" {
		fmt.Printf("  - gemini     ✅  (default: %v)\n", cfg.Gemini.Default)
	}
}

func init() {
	configCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List all configured providers and default")
	configCmd.AddCommand(setCmd)
	rootCmd.AddCommand(configCmd)
}

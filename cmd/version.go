package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/DEEPAKsingh74/zena/internal/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of Zena CLI",
	Long: `Displays the currently installed version of the Zena CLI tool.

This command is helpful to confirm which version you're running,
especially when debugging, reporting issues, or verifying updates.

Example usage:

  zena version

This will output:

  Zena CLI Version: v1.0.0

Note: If you built the binary yourself, the version string may show 'dev'
unless overridden during build using:

  go build -ldflags="-X 'github.com/DEEPAKsingh74/zena/internal/version.Version=v1.0.0'"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Zena CLI Version: %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
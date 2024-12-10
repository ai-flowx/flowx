package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ai-flowx/flowx/config"
)

var rootCmd = &cobra.Command{
	Use:     "flowx",
	Version: config.Version + "-build-" + config.Build,
	Short:   "ai framework",
	Long:    fmt.Sprintf("ai framework %s (%s %s)", config.Version, config.Commit, config.Build),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}
	},
}

// nolint:gochecknoinits
func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

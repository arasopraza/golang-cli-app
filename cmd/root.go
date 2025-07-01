package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bang",
	Short: "bang is a cli tool for get github profile",
	Long:  "bang is a cli tool for get github profile and repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Bang '%s'\n", err)
		os.Exit(1)
	}
}

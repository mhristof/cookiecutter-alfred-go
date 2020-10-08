package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.project_slug}}",
	Short: "{{cookiecutter.project_short_description}}",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute The main function for the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

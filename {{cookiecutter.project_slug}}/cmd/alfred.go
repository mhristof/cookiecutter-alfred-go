package cmd

import (
	"strings"

	"github.com/mhristof/go-alfred"
	"github.com/spf13/cobra"
)

var (
	alfredCmd = &cobra.Command{
		Use:   "alfred",
		Short: "List alfred options",
		Run: func(cmd *cobra.Command, args []string) {
			var opts alfred.ScriptFilter
			if command.Use == "alfred" || command.Use == "help [command]" {
				continue
			}
			//_ = opts.Add(command.Use, command.Use)
			//item := opts.Add(path, path)
			opts.Print()
		},
	}
)

func match(haystack, needle string) string {
	for _, char := range needle {
		haystack = strings.ReplaceAll(haystack, string(char), " ")
	}
	return haystack
}

func init() {
	rootCmd.AddCommand(alfredCmd)
}

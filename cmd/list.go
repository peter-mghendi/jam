package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// listCmd represents the init command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List aliases managed by jam.",
	Long: `List all aliases defined in the .jamrc file.

This command reads your jamfile (~/.jamrc), parses the alias metadata,
and prints a human-friendly list including name, status (enabled/disabled),
and the target command or script.`,
	Example: `  jam list
  jam ls`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.List()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

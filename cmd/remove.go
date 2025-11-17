package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove <name>",
	Aliases: []string{"rm"},
	Short:   "Remove an alias from .jamrc.",
	Long: `Remove an existing alias from your jamfile (~/.jamrc).

This deletes both the alias definition and its associated metadata.
The underlying script or command that the alias pointed to is not touched.`,
	Example: `  # Remove an alias
  jam remove greet

  # Preview changes without writing to ~/.jamrc
  jam remove greet --pretend`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		pretend, err := cmd.Flags().GetBool("pretend")
		cobra.CheckErr(err)

		err = handlers.Remove(name, pretend)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().Bool("pretend", false, "Print the updated .jamrc to STDOUT without writing it to disk.")
}

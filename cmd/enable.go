package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable <name>",
	Short: "Enable an existing alias.",
	Long: `Enable an alias in your jamfile (~/.jamrc).

This marks the alias as enabled in its metadata and ensures the alias
definition is present so it can be used in your shell.`,
	Example: `  # Enable an alias
  jam enable greet

  # Preview changes without writing to ~/.jamrc
  jam enable greet --pretend`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		pretend, err := cmd.Flags().GetBool("pretend")
		cobra.CheckErr(err)

		err = handlers.Toggle(name, true, pretend)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)

	enableCmd.Flags().Bool("pretend", false, "Print the updated .jamrc to STDOUT without writing it to disk.")
}

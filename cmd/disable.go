package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable <name>",
	Short: "Disable an existing alias.",
	Long: `Disable an alias in your jamfile (~/.jamrc).

This marks the alias as disabled in its metadata and ensures the alias
is not active for new shells that source the generated file.`,
	Example: `  # Disable an alias
  jam disable greet

  # Preview changes without writing to ~/.jamrc
  jam disable greet --pretend`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		pretend, err := cmd.Flags().GetBool("pretend")
		cobra.CheckErr(err)

		err = handlers.Toggle(name, false, pretend)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)

	disableCmd.Flags().Bool("pretend", false, "Print the updated .jamrc to STDOUT without writing it to disk.")
}

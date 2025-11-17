package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new .jamrc file.",
	Long: `The init command creates a new, empty .jamrc file.

If a .jamrc file already exists, it will not overwrite it unless the --force flag is provided.
Use --pretend to output the generated file to STDOUT.`,
	Example: `  jam init
  jam init --force
  jam init --pretend`,
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool("force")
		cobra.CheckErr(err)

		pretend, err := cmd.Flags().GetBool("pretend")
		cobra.CheckErr(err)

		err = handlers.Init(force, pretend)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().Bool("force", false, "Overwrite existing .jamrc if it exists.")
	initCmd.Flags().Bool("pretend", false, "Print the generated .jamrc to STDOUT without writing it to disk.")
}

package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// debugCmd represents the init command
var debugCmd = &cobra.Command{
	Use:     "debug",
	Hidden:  true,
	Short:   "Debug the .jamrc file.",
	Long:    "Outputs the AST representation of the .jamrc file to STDOUT.",
	Example: "jam debug",
	Run: func(cmd *cobra.Command, args []string) {
		err := handlers.Debug()
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}

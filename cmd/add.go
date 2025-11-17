package cmd

import (
	"jam/lib/handlers"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <name> <target>",
	Short: "Add a new alias to .jamrc.",
	Long: `Add a new alias to your jamfile (~/.jamrc).

The alias name is what you'll type in your shell, and the target is the
command or script that will be executed. You can optionally attach a
description and start the alias in a disabled state.

By default, aliases are created as enabled.`,
	Example: `  # Simple alias pointing to a script
  jam add greet "$HOME/bin/greet.cs"

  # Alias with a description
  jam add greet "$HOME/bin/greet.cs" --desc "Says hello"

  # Create a disabled alias
  jam add greet "$HOME/bin/greet.cs" --desc "Says hello" --disabled

  # Preview changes without writing to ~/.jamrc
  jam add greet "$HOME/bin/greet.cs" --pretend`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		target := args[1]

		desc, err := cmd.Flags().GetString("desc")
		cobra.CheckErr(err)

		disabled, err := cmd.Flags().GetBool("disabled")
		cobra.CheckErr(err)

		pretend, err := cmd.Flags().GetBool("pretend")
		cobra.CheckErr(err)

		err = handlers.Add(name, target, desc, !disabled, pretend)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("desc", "", "Optional description for the alias.")
	addCmd.Flags().Bool("disabled", false, "Create the alias in a disabled state.")
	addCmd.Flags().Bool("pretend", false, "Print the updated .jamrc to STDOUT without writing it to disk.")
}

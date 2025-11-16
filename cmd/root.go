package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jam",
	Short: "jam: alias manager",
	Long: `jam is a CLI tool for managing your shell aliases through a single declarative file: ~/.jamrc.

It lets you define, organize, and maintain aliases using a structured format, backed by Bash's own
associative arrays. Jam treats .jamrc as the source of truth and updates it safely and consistently
through an AST-aware parser.

With Jam, you can:
- Add, enable, disable, or remove aliases without editing your shell files manually.
- Store alias metadata alongside each alias definition.
- Export aliases to other shells like bash, zsh, or fish.
- Open the underlying command or script in your preferred editor.
- Keep your entire alias ecosystem predictable and version-controllable.`,
	Example: `  # Initialize your system for jam
  jam init 

  # List all managed aliases
  jam list

  # Add a new alias
  jam add greet ~/bin/greet.cs --desc "Hello world"

  # Disable an alias
  jam disable greet

  # Edit an alias command in your editor
  jam edit greet

  # Verify the validity of .jamrc
  jam verify
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

package handlers

import (
	"fmt"
	"jam/lib/jamfile"
	"jam/lib/jamrc"
)

func List() error {
	if exists, err := jamrc.Exists(); err != nil {
		return fmt.Errorf("failed to check if .jamrc exists: %w", err)
	} else if !exists {
		return fmt.Errorf("~/.jamrc does not exist. Use jam init to create one")
	}

	file, err := jamrc.Load()
	if err != nil {
		return err
	}

	document, err := jamfile.FromSyntaxFile(file)
	if err != nil {
		return err
	}

	for _, alias := range document.Aliases {
		fmt.Printf("# %s\n", alias.Description)
		fmt.Printf("# Added at: %s\n", alias.AddedAt)

		if alias.Enabled {
			fmt.Printf("alias %s=\"%s\"\n", alias.Name, alias.Target)
		} else {
			fmt.Printf("# alias %s=\"%s\"\n", alias.Name, alias.Target)
		}
		fmt.Println()
	}

	return nil
}

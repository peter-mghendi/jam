package handlers

import (
	"fmt"
	"jam/lib/jamfile"
	"jam/lib/jamrc"
)

func Remove(name string, pretend bool) error {
	if err := jamrc.Detect(); err != nil {
		return err
	}

	file, err := jamrc.Load()
	if err != nil {
		return err
	}

	document, err := jamfile.FromSyntaxFile(file)
	if err != nil {
		return err
	}

	if _, exists := document.Aliases[name]; !exists {
		return fmt.Errorf("alias %s does not exist", name)
	}

	delete(document.Aliases, name)

	if err := jamfile.Write(document, pretend); err != nil {
		return err
	}

	if !pretend {
		fmt.Printf("[INFO] Removed alias `%s`. Reload your shell to see changes.\n", name)
	}

	return nil
}

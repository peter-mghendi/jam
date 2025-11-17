package handlers

import (
	"fmt"
	"jam/lib/jamfile"
	"jam/lib/jamrc"
	"time"
)

func Add(name, target, description string, enabled, pretend bool) error {
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

	if _, exists := document.Aliases[name]; exists {
		return fmt.Errorf("alias %s already exists", name)
	}

	document.Aliases[name] = jamfile.Alias{
		Name:        name,
		Target:      target,
		Enabled:     enabled,
		Description: description,
		AddedAt:     time.Time{},
	}

	if err := jamfile.Write(document, pretend); err != nil {
		return err
	}

	if !pretend {
		fmt.Printf("[INFO] Added alias `%s`. Reload your shell to see changes.\n", name)
	}

	return nil
}

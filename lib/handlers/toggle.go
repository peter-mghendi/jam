package handlers

import (
	"fmt"
	"jam/lib/jamfile"
	"jam/lib/jamrc"
)

var statuses = map[bool]string{
	true:  "Enabled",
	false: "Disabled",
}

func Toggle(name string, status, pretend bool) error {
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

	alias, exists := document.Aliases[name]
	if !exists {
		return fmt.Errorf("alias %s does not exist", name)
	}

	alias.Enabled = status
	document.Aliases[name] = alias

	if err := jamfile.Write(document, pretend); err != nil {
		return err
	}

	if !pretend {
		fmt.Printf("[INFO] %s alias `%s`. Reload your shell to see changes.\n", statuses[status], name)
	}

	return nil
}

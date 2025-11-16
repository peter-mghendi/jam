package handlers

import (
	"fmt"
	"jam/lib/jamrc"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

func Debug() error {
	if exists, err := jamrc.Exists(); err != nil {
		return fmt.Errorf("failed to check if .jamrc exists: %w", err)
	} else if !exists {
		return fmt.Errorf("~/.jamrc does not exist. Use jam init to create one")
	}

	file, err := jamrc.Load()
	if err != nil {
		return err
	}

	if err = syntax.DebugPrint(os.Stdout, file); err != nil {
		return err
	}

	return nil
}

package handlers

import (
	"fmt"
	"jam/lib/jamfile"
	"jam/lib/jamrc"
)

func Init(force bool, pretend bool) error {
	if exists, err := jamrc.Exists(); err != nil {
		return fmt.Errorf("failed to check if .jamrc exists: %w", err)
	} else if exists && !pretend && force {
		fmt.Printf("[INFO] Overwriting ~/.jamrc. \n")
	} else if exists && !pretend {
		return fmt.Errorf("~/.jamrc already exists. Use --force to overwrite")
	}

	if err := jamfile.Write(&jamfile.Document{}, pretend); err != nil {
		return err
	}

	if !pretend {
		fmt.Printf("[INFO] ~/.jamrc initialized successfully. Add the following lines to the end of ~/.bashrc:\n\n")
		fmt.Println(`  if [ -f ~/.jamrc ]; then
      . ~/.jamrc
  fi`)
	}

	return nil
}

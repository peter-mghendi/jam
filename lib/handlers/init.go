package handlers

import (
	"bytes"
	"fmt"
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

	file := jamrc.Empty()
	buffer := &bytes.Buffer{}
	if err := jamrc.Render(&file, buffer); err != nil {
		return err
	}

	if pretend {
		fmt.Print(buffer.String())
		return nil
	}

	if err := jamrc.Write(buffer); err != nil {
		return err
	}

	fmt.Println("[INFO] ~/.jamrc initialized successfully.")
	return nil
}

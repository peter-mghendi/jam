package jamrc

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Exists returns a boolean representing whether a jamfile exists, or false and an error
func Exists() (bool, error) {
	path, err := jamfile()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Read loads the contents of the jamfile (~/.jamrc) into memory.
func Read() ([]byte, error) {
	path, err := jamfile()
	if err != nil {
		return nil, fmt.Errorf("failed to access .jamrc: %w", err)
	}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf(".jamrc does not exist at %s", path)
	} else if err != nil {
		return nil, fmt.Errorf("failed to read .jamrc: %w", err)
	}

	return data, nil
}

// Write outputs the provided bytes.Buffer to the jamfile dir, or to STDOUT if pretend is true
func Write(buffer *bytes.Buffer) error {
	path, err := jamfile()
	if err != nil {
		return fmt.Errorf("failed to access .jamrc: %w", err)
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("failed to write temp .jamrc: %w", err)
	}

	if err := os.Rename(tmp, path); err != nil {
		return fmt.Errorf("failed to move temp .jamrc into place: %w", err)
	}

	return nil
}

// jamfile returns the path to `~/.jamrc` or an error.
func jamfile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".jamrc"), nil
}

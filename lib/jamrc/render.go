package jamrc

import (
	"bytes"
	"fmt"

	"mvdan.cc/sh/v3/syntax"
)

// Render outputs the provided syntax.File as text into the provided bytes.Buffer
func Render(file *syntax.File, buffer *bytes.Buffer) error {
	printer := syntax.NewPrinter(syntax.Indent(2))
	if err := printer.Print(buffer, file); err != nil {
		return fmt.Errorf("failed to render .jamrc: %w", err)
	}

	if !bytes.HasSuffix(buffer.Bytes(), []byte("\n")) {
		if _, err := buffer.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to finalize .jamrc buffer: %w", err)
		}
	}

	return nil
}

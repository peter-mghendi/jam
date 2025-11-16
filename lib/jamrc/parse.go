package jamrc

import (
	"bytes"
	"fmt"

	"mvdan.cc/sh/v3/syntax"
)

// Load reads ~/.jamrc and parses it into a *syntax.File.
func Load() (*syntax.File, error) {
	data, err := Read()
	if err != nil {
		return nil, err
	}
	return parse(data)
}

// parse parses raw .jamrc contents into a *syntax.File.
func parse(src []byte) (*syntax.File, error) {
	parser := syntax.NewParser(syntax.Variant(syntax.LangBash), syntax.KeepComments(false))
	f, err := parser.Parse(bytes.NewReader(src), ".jamrc")
	if err != nil {
		return nil, fmt.Errorf("failed to parse .jamrc: %w", err)
	}
	return f, nil
}

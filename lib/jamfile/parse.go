package jamfile

import (
	"strings"
	"time"

	"mvdan.cc/sh/v3/syntax"
)

const (
	metaPrefix = "__jam__"

	keyTarget      = "target"
	keyEnabled     = "enabled"
	keyDescription = "description"
	keyAddedAt     = "added_at"
	timeLayout     = time.RFC3339
)

// FromSyntaxFile walks a *syntax.File and builds a Document.
func FromSyntaxFile(f *syntax.File) (*Document, error) {
	document := &Document{Aliases: map[string]Alias{}}

	metaByName := map[string]Alias{}
	targetByName := map[string]string{}

	for _, stmt := range f.Stmts {
		switch x := stmt.Cmd.(type) {

		case *syntax.DeclClause:
			if !isDeclareAssoc(x) {
				continue
			}
			if len(x.Args) < 2 {
				continue
			}

			assign := x.Args[1]
			if assign.Name == nil || assign.Array == nil {
				continue
			}
			rawName := assign.Name.Value // "__jam__greet"
			if !strings.HasPrefix(rawName, metaPrefix) {
				continue
			}
			name := strings.TrimPrefix(rawName, metaPrefix)

			a := metaByName[name]
			a.Name = name

			for _, elem := range assign.Array.Elems {
				key := wordToString(elem.Index.(*syntax.Word))
				val := wordToString(elem.Value)

				switch key {
				case keyTarget:
					a.Target = val
				case keyEnabled:
					a.Enabled = val == "true"
				case keyDescription:
					a.Description = val
				case keyAddedAt:
					if t, err := time.Parse(timeLayout, val); err == nil {
						a.AddedAt = t
					}
				}
			}

			metaByName[name] = a

		case *syntax.CallExpr:
			if len(x.Args) < 2 {
				continue
			}

			if !wordIsLiteral(x.Args[0], "alias") {
				continue
			}

			name, target := parseAliasWord(x.Args[1])
			if name == "" {
				continue
			}
			targetByName[name] = target
		}
	}

	for name, meta := range metaByName {
		meta.Target = coalesce(meta.Target, targetByName[name])
		document.Aliases[name] = meta
	}

	return document, nil
}

func isDeclareAssoc(d *syntax.DeclClause) bool {
	if d.Variant.Value != "declare" {
		return false
	}
	if len(d.Args) == 0 {
		return false
	}

	first := d.Args[0]
	if first.Value == nil {
		return false
	}
	return wordToString(first.Value) == "-A"
}

// wordToString flattens a *syntax.Word into a string, handling
// literals, single quotes, double quotes, and simple parameter expansions.
func wordToString(w *syntax.Word) string {
	if w == nil {
		return ""
	}

	var b strings.Builder
	for _, p := range w.Parts {
		switch v := p.(type) {
		case *syntax.Lit:
			b.WriteString(v.Value)
		case *syntax.SglQuoted:
			b.WriteString(v.Value)
		case *syntax.DblQuoted:
			for _, qp := range v.Parts {
				switch qv := qp.(type) {
				case *syntax.Lit:
					b.WriteString(qv.Value)
				case *syntax.ParamExp:
					// Using $HOME for tilde.
					b.WriteString("$")
					b.WriteString(qv.Param.Value)
				}
			}
		}
	}
	return b.String()
}

// wordIsLiteral checks if the word is a single literal with the given value.
func wordIsLiteral(w *syntax.Word, value string) bool {
	if w == nil || len(w.Parts) != 1 {
		return false
	}
	lit, ok := w.Parts[0].(*syntax.Lit)
	return ok && lit.Value == value
}

// parseAliasWord extracts name and target from a syntax.Word
func parseAliasWord(w *syntax.Word) (name, target string) {
	if w == nil || len(w.Parts) == 0 {
		return "", ""
	}

	// First part must be a literal of form "name="
	firstLit, ok := w.Parts[0].(*syntax.Lit)
	if !ok {
		return "", ""
	}
	raw := firstLit.Value // e.g. "greet="
	if !strings.Contains(raw, "=") {
		return "", ""
	}
	name = strings.SplitN(raw, "=", 2)[0]

	// The target is the rest of the word; easiest is to rebuild from all parts,
	// but discard the "name=" prefix.
	full := wordToString(w)
	prefix := name + "="
	if strings.HasPrefix(full, prefix) {
		target = strings.TrimPrefix(full, prefix)
	} else {
		target = full
	}

	return name, target
}

func coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

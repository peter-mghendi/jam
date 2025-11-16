package jamfile

import "time"

type Document struct {
	Aliases []Alias
}

type Alias struct {
	Name        string
	Target      string
	Enabled     bool
	Description string
	AddedAt     time.Time
}

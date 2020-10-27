package objectstore

import "time"

type Stat struct {
	Name         string
	IsFolder     bool
	Size         int64
	ModifiedTime time.Time
}

type Stats []*Stat

// Len is the number of elements in the collection.
func (s Stats) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s Stats) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap swaps the elements with indexes i and j.
func (s Stats) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

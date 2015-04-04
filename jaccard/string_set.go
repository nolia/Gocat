package jaccard

import "bytes"

type StringSet map[string]bool

func NewSet() StringSet {
	m := make(map[string]bool)
	return StringSet(m)
}

func NewSetWithValues(arr []string) StringSet {
	set := NewSet()
	for _, v := range arr {
		set.Add(v)
	}
	return set
}

// Adds and element to set. If set contained it before - returns false.
func (set StringSet) Add(s string) bool {
	if _, ok := set[s]; ok {
		return false
	}
	set[s] = true

	return true
}

// Returns true if set contains an element, and false otherwise.
func (set StringSet) Has(s string) bool {
	_, ok := set[s]
	return ok
}

// Removes an element from set. Returns true if it was removed, and false otherwise.
func (set StringSet) Remove(s string) bool {
	if !set.Has(s) {
		return false
	}
	delete(set, s)
	return true
}

func (set StringSet) String() string {
	var buf bytes.Buffer
	comma := false

	buf.WriteRune('{')
	for v, _ := range set {
		if comma {
			buf.WriteString(", ")
		}
		comma = true
		buf.WriteString(v)
	}
	buf.WriteRune('}')

	return buf.String()
}

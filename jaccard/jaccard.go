package jaccard

import (
	"bytes"
	"regexp"
)

type StringSet map[string]bool

func NewSet() StringSet {
	m := make(map[string]bool)
	return StringSet(m)
}

func (set StringSet) Add(s string) bool {
	if _, ok := set[s]; ok {
		return false
	}
	set[s] = true

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

func SplitWords(s string) []string {
	re := regexp.MustCompile("(\\w+)")
	return re.FindAllString(s, -1)
}

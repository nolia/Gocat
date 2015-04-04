package jaccard

import "regexp"

func SplitWords(s string) []string {
	re := regexp.MustCompile("(\\w+)")
	return re.FindAllString(s, -1)
}

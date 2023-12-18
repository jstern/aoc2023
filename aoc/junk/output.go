package junk

import (
	"strings"

	"github.com/samber/lo"
)

// SideBySide takes a separator and set of multiline strings
// and returns a string with lines for each joined by the separator.
func SideBySide(sep string, items ...string) string {
	res := make([]string, 0)
	itemRows := lo.Map(items, func(item string, _ int) []string {
		return strings.Split(item, "\n")
	})
	for row := range strings.Split(items[0], "\n") {
		v := make([]string, 0)
		for i := range items {
			v = append(v, itemRows[i][row])
		}
		res = append(res, strings.Join(v, sep))
	}
	return strings.Join(res, "\n")
}

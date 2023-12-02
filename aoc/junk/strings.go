package junk

import (
	"strings"

	"github.com/samber/lo"
)

// ReverseString reverses a string.
func ReverseString(s string) string {
	return strings.Join(lo.Reverse(strings.Split(s, "")), "")
}

package solution

import (
	"sort"
)

var solutions = make(map[string]Solution)

// A Solution is just a function that takes an input string and returns an answer string.
type Solution func(inp string) (answer string)

// For returns the registered Solution for the supplied key.
func For(key string) Solution {
	return solutions[key]
}

func register(key string, soln Solution) {
	solutions[key] = soln
}

// List returns a sorted list of registerd solution keys.
func List() []string {
	keys := make([]string, 0, len(solutions))
	for k := range solutions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

package aoc

import (
	"sort"
)

var solutions = make(map[string]Solution)

// A Solution is just a function that takes an input string and returns an answer string.
type Solution func(inp string) (answer string)

// SolutionFor returns the registered Solution for the supplied key.
func SolutionFor(key string) Solution {
	return solutions[key]
}

func registerSolution(key string, soln Solution) {
	solutions[key] = soln
}

// ListSolutions returns a sorted list of registered solution keys.
func ListSolutions() []string {
	keys := make([]string, 0, len(solutions))
	for k := range solutions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

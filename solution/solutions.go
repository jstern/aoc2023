package solution

import "fmt"

var solutions = make(map[string]Solution)

type Solution func(inp string) (answer string)

func For(key string) Solution {
	return solutions[key]
}

func register(key string, soln Solution) {
	solutions[key] = soln
}

func List() {
	fmt.Println("Solutions registered:")
	for k, fn := range solutions {
		fmt.Printf(" * %s %v\n", k, fn)
	}
}

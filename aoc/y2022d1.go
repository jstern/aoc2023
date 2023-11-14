package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2022:1:1", y2022d1part1)
	registerSolution("2022:1:2", y2022d1part2)

}

func y2022d1part1(input string) string {
	max := 0
	cur := 0
	for _, str := range strings.Split(input+"\n", "\n") {
		cals, err := strconv.Atoi(str)
		if err != nil {
			// assume empty line, check for new max
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}
		cur += cals
	}
	return fmt.Sprintf("%d", max)
}

func y2022d1part2(input string) string {
	top := make([]int, 3)
	cur := 0
	for _, str := range strings.Split(input+"\n", "\n") {
		cals, err := strconv.Atoi(str)
		if err != nil {
			top = append(top, cur)
			slices.Sort(top) // reorder
			top = top[1:]    // drop lowest
			cur = 0
			continue
		}
		cur += cals
	}
	return fmt.Sprintf("%d", lo.Sum(top))
}

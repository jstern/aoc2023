package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:2:1", y2023d2part1)
	registerSolution("2023:2:2", y2023d2part2)

	cubeGameRe = regexp.MustCompile(`^Game ([0-9]+): (.*)?`)
}

var cubeGameRe *regexp.Regexp

func y2023d2part1(input string) string {
	tot := 0
	for _, desc := range strings.Split(input, "\n") {
		if desc == "" {
			continue
		}
		game := parseCubeGame(desc)
		if game.possibleFor(cubeTotals) {
			tot += game.id
		}
	}
	return fmt.Sprintf("%d", tot)
}

var cubeTotals = map[string]int{"red": 12, "green": 13, "blue": 14}

type cubeGame struct {
	id      int
	maxPull map[string]int
}

func parseCubeGame(desc string) cubeGame {
	// the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then *put them back* in the bag
	g := cubeGame{maxPull: make(map[string]int)}
	details := cubeGameRe.FindStringSubmatch(desc)
	g.id, _ = strconv.Atoi(details[1])

	for _, pull := range strings.Split(details[2], "; ") {
		for _, c := range strings.Split(pull, ", ") {
			cc := strings.Split(c, " ")
			color := cc[1]
			count, _ := strconv.Atoi(cc[0])
			if count > g.maxPull[color] {
				g.maxPull[color] = count
			}
		}
	}
	return g
}

func (cg cubeGame) possibleFor(maxes map[string]int) bool {
	return !lo.Contains(
		lo.MapToSlice(cg.maxPull, func(c string, m int) bool {
			return maxes[c] >= m
		}),
		false,
	)
}

func (cg cubeGame) power() int {
	// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
	// maxPull should give us "the fewest number of cubes of each color"
	return lo.Reduce(lo.Values(cg.maxPull), func(a int, m int, _ int) int { return a * m }, 1)
}

func y2023d2part2(input string) string {
	tot := 0
	for _, desc := range strings.Split(input, "\n") {
		if desc == "" {
			continue
		}
		game := parseCubeGame(desc)
		tot += game.power()
	}
	return fmt.Sprintf("%d", tot)
}

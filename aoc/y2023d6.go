package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:6:1", y2023d6part1)
	registerSolution("2023:6:2", y2023d6part2)
}

func y2023d6part1(input string) string {
	races := parseBoatRaces(strings.Split(input, "\n"))
	wins := lo.Map(races, func(br lo.Tuple2[int, int], _ int) int { return boatRaceWins(br) })
	return fmt.Sprintf("%d", lo.Reduce(wins, func(a int, t int, _ int) int { return a * t }, 1))
}

func parseBoatRaces(lines []string) []lo.Tuple2[int, int] {
	times := lo.Map(strings.Fields(lines[0])[1:], func(s string, _ int) int { t, _ := strconv.Atoi(s); return t })
	recs := lo.Map(strings.Fields(lines[1])[1:], func(s string, _ int) int { t, _ := strconv.Atoi(s); return t })
	return lo.Zip2(times, recs)
}

func boatRaceWins(br lo.Tuple2[int, int]) int {
	tot := 0
	for press := 1; press < br.A; press++ {
		dist := press * (br.A - press)
		if dist > br.B {
			tot += 1
		}
	}
	return tot
}

func y2023d6part2(input string) string {
	time, record := parseBoatRaces2(strings.Split(input, "\n"))

	n := boatRaceWins(lo.Tuple2[int, int]{A: time, B: record})

	return fmt.Sprintf("%d", n)
}

func parseBoatRaces2(lines []string) (int, int) {
	time, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	dist, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	return time, dist
}

package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registerSolution("2023:1:1", y2023d1part1)
	registerSolution("2023:1:2", y2023d1part2)
}

func y2023d1part1(input string) string {
	tot := 0
	for _, line := range strings.Split(input, "\n") {
		first := ""
		last := ""
		for _, c := range strings.Split(line, "") {
			if _, err := strconv.Atoi(c); err == nil {
				if first == "" {
					first = c
				}
				last = c
			}
		}
		if lt, err := strconv.Atoi(first + last); err == nil {
			tot += lt
		}
	}
	return fmt.Sprintf("%d", tot)
}

func y2023d1part2(input string) string {
	tot := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		v := calibrationValue(line)
		tot += v
	}
	return fmt.Sprintf("%d", tot)
}

var numberWords = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func calibrationValue(line string) int {
	lo := len(line)
	d1 := ""
	hi := -1
	d2 := ""
	for n, v := range numberWords {
		ni := strings.Index(line, n)
		if ni >= 0 {
			lni := strings.LastIndex(line, n)
			if lni > hi {
				hi = lni
				d2 = v
			}
			if ni <= lo {
				lo = ni
				d1 = v
			}
		}
		vi := strings.Index(line, v)
		if vi >= 0 {
			lvi := strings.LastIndex(line, v)
			if lvi > hi {
				hi = lvi
				d2 = v
			}
			if vi <= lo {
				lo = vi
				d1 = v
			}
		}
	}
	res, err := strconv.Atoi(d1 + d2)
	if err != nil {
		panic(err)
	}
	return res
}

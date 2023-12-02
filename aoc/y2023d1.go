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

type calibrationCalc struct {
	lo int
	d1 string
	hi int
	d2 string
}

func (c *calibrationCalc) update(first, last int, val string) {
	if first < 0 {
		return
	}
	if last > c.hi {
		c.hi = last
		c.d2 = val
	}
	if first <= c.lo {
		c.lo = first
		c.d1 = val
	}
}

func (c calibrationCalc) result() int {
	res, err := strconv.Atoi(c.d1 + c.d2)
	if err != nil {
		panic(err)
	}
	return res
}

func calibrationValue(line string) int {
	calc := &calibrationCalc{lo: len(line), hi: -1}
	for n, v := range numberWords {
		ni := strings.Index(line, n)
		lni := strings.LastIndex(line, n)
		calc.update(ni, lni, v)

		vi := strings.Index(line, v)
		lvi := strings.LastIndex(line, v)
		calc.update(vi, lvi, v)
	}
	return calc.result()
}

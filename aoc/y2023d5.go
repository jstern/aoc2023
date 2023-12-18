package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:5:1", y2023d5part1)
	registerSolution("2023:5:2", y2023d5part2)
}

func y2023d5part1(input string) string {
	lines := strings.Split(input, "\n")
	seeds := parseSeeds(lines[0])

	var mapper gardenMapper
	for _, line := range lines[2:] {
		if strings.HasSuffix(line, "map:") {
			mapper = gardenMapper{ranges: make(map[int]gardenRange)}
		} else if len(line) > 0 {
			gmr := parseRange(line)
			mapper.ranges[gmr.start] = gmr
		} else {
			applyMapper(mapper, seeds)
		}
	}
	return fmt.Sprintf("%d", lo.Min(lo.Values(seeds)))
}

func parseSeeds(line string) map[int]int {
	seeds := make(map[int]int)
	for _, n := range strings.Fields(line) {
		s, err := strconv.Atoi(n)
		if err == nil {
			seeds[s] = s
		}
	}
	return seeds
}

func applyMapper(gm gardenMapper, seeds map[int]int) {
	updated := junk.NewSet[int]()
	rss := lo.Keys(gm.ranges)
	slices.Sort(rss)
	for _, rs := range rss {
		gmr := gm.ranges[rs]
		for seed, dst := range seeds {
			if updated.Contains(seed) {
				continue
			}
			newdst, moved := gmr.translate(dst)
			if moved {
				updated.Add(seed)
				seeds[seed] = newdst
			}
		}
	}
}

type gardenMapper struct {
	ranges map[int]gardenRange // ranges keys by source start
}

type gardenRange struct {
	start int
	delta int
	max   int
}

func (gr gardenRange) String() string {
	pn := ""
	if gr.delta >= 0 {
		pn = "+"
	}
	return fmt.Sprintf("[%d-%d]:%s%d", gr.start, gr.max, pn, gr.delta)
}

func parseRange(line string) gardenRange {
	parts := lo.Map(strings.Fields(line), func(s string, _ int) int { n, _ := strconv.Atoi(s); return n })
	return gardenRange{
		start: parts[1],
		delta: parts[0] - parts[1],
		max:   parts[1] + parts[2],
	}
}

func (r gardenRange) translate(dst int) (int, bool) {
	if dst < r.start || dst >= r.max {
		return dst, false
	}
	return dst + r.delta, true
}

func y2023d5part2(input string) string {
	lines := strings.Split(input, "\n")
	seeds := parseSeedRanges(lines[0])
	fmt.Println(seeds)

	var mapper gardenMapper
	for _, line := range lines[2:] {
		if strings.HasSuffix(line, "map:") {
			fmt.Println("\n" + line)
			mapper = gardenMapper{ranges: make(map[int]gardenRange)}
		} else if len(line) > 0 {
			gmr := parseRange(line)
			mapper.ranges[gmr.start] = gmr
		} else {
			fmt.Println("applying to", seeds)
			seeds = applyMapper2(mapper, seeds)
			fmt.Println("result", seeds)
		}
	}
	return fmt.Sprintf("%d", lo.Min(lo.Keys(seeds)))
}

func parseSeedRanges(line string) map[int]gardenRange {
	seeds := make(map[int]gardenRange)
	start := -1
	for _, n := range strings.Fields(line) {
		s, err := strconv.Atoi(n)
		if err == nil {
			if start == -1 {
				start = s
			} else {
				seeds[start] = gardenRange{start: start, max: start + s - 1}
				start = -1
			}
		}
	}
	return seeds
}

func (r gardenRange) translate2(sr gardenRange) (processed []gardenRange, remainder *gardenRange) {
	processed = make([]gardenRange, 0)

	// if all of sr is after our range, just return it as the remainder
	if sr.start > r.max {
		return processed, &sr
	}

	// if all of sr is before our range, consider it processed with no remainder
	if sr.max < r.start {
		return append(processed, sr), nil
	}

	// if some of sr is in our range, the remainder is any part that is in sr but after the end of our range
	if sr.max > r.max {
		remainder = &gardenRange{start: r.max + 1, max: sr.max}
	}

	// first, consider any part of the input range before our range to be processed as-is
	if sr.start < r.start {
		maxBefore := lo.Min([]int{sr.max, r.start - 1})
		processed = append(processed, gardenRange{start: sr.start, max: maxBefore})
	}

	// finally, process any part of the input range overlapping with our range to
	// if we got here, we know !(sr.max < r.start)
	max := lo.Min([]int{r.max, sr.max})
	processed = append(processed, gardenRange{start: r.start + r.delta, max: max + r.delta})

	return processed, remainder
}

func applyMapper2(gm gardenMapper, seeds map[int]gardenRange) map[int]gardenRange {
	result := make([]gardenRange, 0)

	seedRangeStarts := lo.Keys(seeds)
	slices.Sort(seedRangeStarts)

	mapRangeStarts := lo.Keys(gm.ranges)
	slices.Sort(mapRangeStarts)

	for _, srs := range seedRangeStarts {
		seedRange := seeds[srs]
		for _, mrs := range mapRangeStarts {
			mapRange := gm.ranges[mrs]
			processed, remainder := mapRange.translate2(seedRange)
			fmt.Printf("%s x %s -> %s | %s\n", mapRange, seedRange, processed, remainder)
			result = append(result, processed...)
			if remainder != nil {
				seedRange = *remainder
			}
		}
	}
	return lo.Associate(
		// remove empty ranges
		lo.Filter(result, func(r gardenRange, _ int) bool { return r.max >= r.start }),
		func(r gardenRange) (int, gardenRange) { return r.start, r },
	)
}

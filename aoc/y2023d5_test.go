package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d5Input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`

func Test_2023_Day_5_Part_1_Example(t *testing.T) {
	result := y2023d5part1(y2023d5Input)
	assert.Equal(t, "35", result)
}

func Test_2023_Day_5_Part_1_ApplyMapper(t *testing.T) {
	seeds := map[int]int{79: 79, 14: 14, 55: 55, 13: 13}
	mapper := gardenMapper{ranges: map[int]gardenRange{
		98: parseRange("50 98 2"),
		50: parseRange("52 50 48"),
	}}
	applyMapper(mapper, seeds)
	assert.Equal(t, map[int]int{79: 81, 14: 14, 55: 57, 13: 13}, seeds)
}

func Test_2023_Day_5_Part_2_Example(t *testing.T) {
	result := y2023d5part2(y2023d5Input)
	assert.Equal(t, "46", result)
	t.Fail()
}

func Test_2023_Day_5_translate2(t *testing.T) {
	sr := gardenRange{start: 0, max: 3}

	// [4-5] x [0-3] -> [0-3] | nil
	processed, remainder := gardenRange{start: 4, max: 5}.translate2(sr)
	assert.Nil(t, remainder)
	assert.Equal(t, []gardenRange{{start: 0, max: 3}}, processed)

	// [0-3](+10) x [0-3] -> [10-13] | nil
	processed, remainder = gardenRange{start: 0, max: 3, delta: 10}.translate2(sr)
	assert.Nil(t, remainder)
	assert.Equal(t, []gardenRange{{start: 10, max: 13}}, processed)

	// [2-5](+10) x [0-3] -> [0-1], [12-13] | nil
	processed, remainder = gardenRange{start: 2, max: 5, delta: 10}.translate2(sr)
	assert.Nil(t, remainder)
	assert.Equal(t, []gardenRange{{start: 0, max: 1}, {start: 12, max: 13}}, processed)

	// [1-2](+10) x [0-3] -> [0-0], [11-12] | [3-3]
	processed, remainder = gardenRange{start: 1, max: 2, delta: 10}.translate2(sr)
	assert.Equal(t, gardenRange{start: 3, max: 3}, *remainder)
	assert.Equal(t, []gardenRange{{start: 0, max: 0}, {start: 11, max: 12}}, processed)
}

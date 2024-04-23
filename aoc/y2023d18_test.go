package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d18Input = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`

func Test_2023_Day_18_Part_1_Example(t *testing.T) {
	result := y2023d18part1(y2023d18Input)
	assert.Equal(t, "62", result)
}

func Test_2023_Day_18_Part_1_excavateRow(t *testing.T) {
	row := []bool{false, false, true, false, true, false, false, true, false, true, false, false}
	rs := rowstr(row)
	assert.Equal(t, `..#.#..#.#..`, rs)
	excavateRow(row, rs)
	assert.Equal(t, `..###..###..`, rowstr(row))

	row = []bool{false, false, true, false, true, true, true, false, false, true, false, true, false, false}
	rs = rowstr(row)
	assert.Equal(t, `..#.###..#.#..`, rs)
	excavateRow(row, rs)
	assert.Equal(t, `..#####..###..`, rowstr(row))

	// ..########.#..####....................########..................................#.....#...
	//`..#.#..#.#..#.#..######..`
}

func Test_2023_Day_18_Part_2_Example(t *testing.T) {
	//result := y2023d18part2(y2023d18Input)
	//assert.Equal(t, "still right!", result)
	t.Skip()
}

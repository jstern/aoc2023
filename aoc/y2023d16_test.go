package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d16Input = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`

func Test_2023_Day_16_Part_1_Example(t *testing.T) {
	result := y2023d16part1(y2023d16Input)
	assert.Equal(t, "46", result)
}

func Test_2023_Day_16_Part_2_Example(t *testing.T) {
	//result := y2023d16part2(y2023d16Input)
	//assert.Equal(t, "still right!", result)
	t.Skip()
}

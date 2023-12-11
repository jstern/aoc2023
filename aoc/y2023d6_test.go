package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d6Input = `Time:      7  15   30
Distance:  9  40  200
`

func Test_2023_Day_6_Part_1_Example(t *testing.T) {
	result := y2023d6part1(y2023d6Input)
	assert.Equal(t, "288", result)
}

func Test_2023_Day_6_Part_2_Example(t *testing.T) {
	//result := y2023d6part2(y2023d6Input)
	//assert.Equal(t, "still right!", result)
	t.Skip()
}

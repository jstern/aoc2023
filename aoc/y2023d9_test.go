package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

var y2023d9Input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func Test_2023_Day_9_Part_1_Example(t *testing.T) {
	result := y2023d9part1(y2023d9Input)
	assert.Equal(t, "114", result)
}

func Test_2023_Day_9_Part_2_Example(t *testing.T) {
	result := y2023d9part2(y2023d9Input)
	assert.Equal(t, "2", result)
}

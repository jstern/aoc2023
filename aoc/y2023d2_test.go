package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d2partInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func Test_2023_Day_2_Part_1_Example(t *testing.T) {
	result := y2023d2part1(y2023d2partInput)
	assert.Equal(t, "8", result)
}

func Test_2023_Day_2_Part_2_Example(t *testing.T) {
	result := y2023d2part2(y2023d2partInput)
	assert.Equal(t, "2286", result)
}

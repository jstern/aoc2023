package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d8Input = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

var y2023d8Input2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func Test_2023_Day_8_Part_1_Examples(t *testing.T) {
	result := y2023d8part1(y2023d8Input)
	assert.Equal(t, "2", result)
	result2 := y2023d8part1(y2023d8Input2)
	assert.Equal(t, "6", result2)
}

var y2023d8part2Input = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

func Test_2023_Day_8_Part_2_Example(t *testing.T) {
	result := y2023d8part2(y2023d8part2Input)
	assert.Equal(t, "6", result)
}

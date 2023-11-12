package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2022d1partInput = `xxx
	yyy`

func Test_2022_Day_1_Part_1_Example(t *testing.T) {
	result := y2022d1part1(y2022d1partInput)
	assert.Equal(t, "right", result)
}

func Test_2022_Day_1_Part_2_Example(t *testing.T) {
	result := y2022d1part2(y2022d1partInput)
	assert.Equal(t, "still right!", result)
}

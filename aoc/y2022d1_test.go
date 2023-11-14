package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2022d1partInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test_2022_Day_01_Part_1_Example(t *testing.T) {
	result := y2022d1part1(y2022d1partInput)
	assert.Equal(t, "24000", result)
}

func Test_2022_Day_01_Part_2_Example(t *testing.T) {
	result := y2022d1part2(y2022d1partInput)
	assert.Equal(t, "45000", result)
}

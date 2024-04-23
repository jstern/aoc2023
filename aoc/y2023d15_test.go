package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d15Input = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func Test_2023_Day_15_Part_1_Example(t *testing.T) {
	result := y2023d15part1(y2023d15Input)
	assert.Equal(t, "1320", result)
}

func Test_2023_Day_15_Part_2_Example(t *testing.T) {
	result := y2023d15part2(y2023d15Input)
	assert.Equal(t, "145", result)
}

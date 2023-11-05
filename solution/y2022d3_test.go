package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2022d3partInput = `xxx
	yyy`

func Test_y2022d3part1(t *testing.T) {
	result := y2022d3part1(y2022d3partInput)
	assert.Equal(t, "right", result)
}

func Test_y2022d3part2(t *testing.T) {
	result := y2022d3part2(y2022d3partInput)
	assert.Equal(t, "still right!", result)
}

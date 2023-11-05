package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_y2022d1part1(t *testing.T) {
	inp := `1000
	2000
	3000
	
	4000
	
	5000
	6000
	
	7000
	8000
	9000
	
	10000`

	result := y2022d1part1(inp)

	assert.Equal(t, "wrong", result)
}

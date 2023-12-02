package aoc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d1partInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func Test_2023_Day_1_Part_1_Example(t *testing.T) {
	result := y2023d1part1(y2023d1partInput)
	assert.Equal(t, "142", result)
}

var y2023d1part2Input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
eightwo
eighthree
oneight
nineight
twone
2mqfhbpn
`

func Test_2023_Day_1_Part_2_Example(t *testing.T) {
	result := y2023d1part2(y2023d1part2Input)
	assert.Equal(t, "605", result)
}

func Test_2023_Day_1_Part_2_calibrationValue(t *testing.T) {
	type tc struct {
		line string
		want int
	}
	lines := strings.Split(y2023d1part2Input+"\n", "\n")
	tests := make([]tc, 13)
	for i, w := range []int{29, 83, 13, 24, 42, 14, 76, 82, 83, 18, 98, 21, 22} {
		tests[i] = tc{line: lines[i], want: w}
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, calibrationValue(tt.line), tt.line)
	}
}

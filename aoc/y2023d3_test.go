package aoc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d3partInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func Test_2023_Day_3_Part_1_Example(t *testing.T) {
	result := y2023d3part1(y2023d3partInput)
	assert.Equal(t, "4361", result)
}

func Test_2023_Day_3_Part_1_Goof_1(t *testing.T) {
	// seeing numbers followed by symbols not being counted!
	result := y2023d3part1(".2+..")
	assert.Equal(t, "2", result)
}

var goof2 = strings.TrimSpace(`
.....664...998........343...............851............................2............414.....................3....................948.164....
......*..................*617....885...*....................-......250.........536..........470...#..................../4......=.....*......
`)

func Test_2023_Day_3_Part_1_Goof_2(t *testing.T) {
	es := parseSchematic(goof2)
	fmt.Println(es.symbols)
	fmt.Println(es.nums)
	assert.Equal(t, 2643, es.partsSum)
}

func Test_2023_Day_3_Part_2_Example(t *testing.T) {
	result := y2023d3part2(y2023d3partInput)
	assert.Equal(t, "467835", result)
}

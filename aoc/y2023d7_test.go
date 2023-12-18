package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2023d7Input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func Test_2023_Day_7_Part_1_Example(t *testing.T) {
	result := y2023d7part1(y2023d7Input)
	assert.Equal(t, "6440", result)
}

func Test_2023_Day_7_handType(t *testing.T) {
	tests := []struct {
		cards []string
		want  string
	}{
		{[]string{"A", "A", "A", "A", "A"}, "5k"},
		{[]string{"A", "A", "Q", "A", "A"}, "4k"},
		{[]string{"A", "A", "Q", "A", "Q"}, "FH"},
		{[]string{"A", "A", "4", "A", "Q"}, "3k"},
		{[]string{"A", "A", "4", "4", "Q"}, "2p"},
		{[]string{"A", "A", "4", "3", "Q"}, "2k"},
		{[]string{"2", "9", "4", "7", "Q"}, "HC"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := handType(tt.cards); got != tt.want {
				t.Errorf("handType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_2023_Day_7_compareHands(t *testing.T) {
	tests := []struct {
		h1   string
		h2   string
		want bool
	}{
		{"KTJJT 220", "QQQJA 483", true},
	}
	for _, tt := range tests {
		t.Run(tt.h1+" vs "+tt.h2, func(t *testing.T) {
			if got := compareHands(parseHand(tt.h1, false), parseHand(tt.h2, false), camelValues); got != tt.want {
				t.Errorf("compareHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_2023_Day_7_Part_2_Example(t *testing.T) {
	result := y2023d7part2(y2023d7Input)
	assert.Equal(t, "5905", result)
}

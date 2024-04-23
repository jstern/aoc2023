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
	result = y2023d7part1(y2023d7AltInputFromReddit)
	assert.Equal(t, "6592", result)
}

func Test_2023_Day_7_handType(t *testing.T) {
	tests := []struct {
		cards string
		want  string
	}{
		{"AAAAA", "5k"},
		{"AAQAA", "4k"},
		{"AAQAQ", "FH"},
		{"AA4AQ", "3k"},
		{"AA44Q", "2p"},
		{"AA43Q", "2k"},
		{"2947Q", "HC"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := handType(tt.cards); got != tt.want {
				t.Errorf("handType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_2023_Day_7_bestHandType(t *testing.T) {
	tests := []struct {
		cards string
		ht    string
	}{
		{"2QK4A", "HC"},
		{"JQ4KA", "2k"},
		{"QQ4KA", "2k"},
		{"234JJ", "3k"},
		{"Q2Q3J", "3k"},
		{"Q2JKK", "3k"},
		{"232JJ", "4k"},
		{"555J6", "4k"},
		{"55JJ6", "4k"},
		{"J4JJJ", "5k"},
		{"44JJJ", "5k"},
		{"2233J", "FH"},
		{"26J77", "3k"},
		{"7JJJ7", "5k"},
		{"JJ88K", "4k"},
	}
	for _, tt := range tests {
		t.Run(tt.ht, func(t *testing.T) {
			ht, _ := bestHandType(tt.cards)
			assert.Equal(t, tt.ht, ht)
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

var y2023d7AltInputFromReddit = `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41
`

var y2023d7Alt2 = ` 2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JKQKK 21
JJJJ2 41
`

func Test_2023_Day_7_Part_2_Example(t *testing.T) {
	result := y2023d7part2(y2023d7Input)
	assert.Equal(t, "5905", result)
	result = y2023d7part2(y2023d7AltInputFromReddit)
	assert.Equal(t, "6839", result)
	result = y2023d7part2(y2023d7Alt2)
	assert.Equal(t, "7460", result)
}

func Test_2023_Day_7_Part_2_substitutions(t *testing.T) {
	choices := []rune{'A', 'B', 'C'}
	substs := substitutions("", 3, choices)
	assert.Len(t, substs, 27)
}

func Test_2023_Day_7_compareHandsWithJokers(t *testing.T) {
	tests := []struct {
		h1   string
		h2   string
		want bool
	}{
		{"88J88", "9J99J", true},
	}
	for _, tt := range tests {
		t.Run(tt.h1+" vs "+tt.h2, func(t *testing.T) {
			if got := compareHands(camelHand{cards: tt.h1, typ: "5k"}, camelHand{cards: tt.h2, typ: "5k"}, jokerValues); got != tt.want {
				t.Errorf("compareHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

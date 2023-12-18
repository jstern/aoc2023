package aoc

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:7:1", y2023d7part1)
	registerSolution("2023:7:2", y2023d7part2)
}

var camelValues = map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}
var jokerValues = map[string]int{"A": 13, "K": 12, "Q": 11, "T": 10, "9": 8, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1}

var camelStrengths = map[string]int{"5k": 6, "4k": 5, "FH": 4, "3k": 3, "2p": 2, "2k": 1, "HC": 0}

type camelHand struct {
	cards []string
	typ   string
	bid   int
}

func compareHands(h1, h2 camelHand, values map[string]int) bool {
	// returns true if h2 is higher rank than h1
	h2s := camelStrengths[h2.typ]
	h1s := camelStrengths[h1.typ]
	if h2s != h1s {
		return h2s > h1s
	}
	for i, h2c := range h2.cards {
		h2ci := values[h2c]
		h1ci := values[h1.cards[i]]
		if h2ci != h1ci {
			return h2ci > h1ci
		}
	}
	return true
}

func handType(cards []string) string {
	counter := make(map[string]int)
	for _, c := range cards {
		counter[c] = counter[c] + 1
	}
	counts := lo.Values(counter)
	slices.Sort(counts)
	slices.Reverse(counts)
	hiCount := lo.Max(counts)
	switch hiCount {
	case 5:
		return "5k"
	case 4:
		return "4k"
	case 3:
		if lo.Contains(counts, 2) {
			return "FH"
		}
		return "3k"
	case 2:
		if counts[1] == 2 {
			return "2p"
		}
		return "2k"
	case 1:
		return "HC"
	}
	panic("can't figure out hand type")
}

func bestHandType(cards []string) string {
	//baseHT := handType(lo.Filter[](cards, ))
	// todo: determine best hand type with jokers
	return handType(cards)
}

func parseHand(desc string, jokers bool) camelHand {
	parts := strings.Fields(desc)
	cards := strings.Split(parts[0], "")
	bid, _ := strconv.Atoi(parts[1])
	ht := handType(cards)
	if jokers && lo.Contains(cards, "J") {
		ht = bestHandType(cards)
	}
	return camelHand{
		cards: cards,
		typ:   ht,
		bid:   bid,
	}
}

func y2023d7part1(input string) string {
	hands := lo.Map(
		strings.Split(strings.TrimSuffix(input, "\n"), "\n"),
		func(line string, _ int) camelHand {
			return parseHand(line, false)
		},
	)
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], camelValues)
	})
	res := lo.Reduce(hands, func(a int, h camelHand, i int) int { return a + ((i + 1) * h.bid) }, 0)
	return fmt.Sprintf("%d", res)
}

func y2023d7part2(input string) string {
	hands := lo.Map(
		strings.Split(strings.TrimSuffix(input, "\n"), "\n"),
		func(line string, _ int) camelHand {
			return parseHand(line, true)
		},
	)
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], jokerValues)
	})
	res := lo.Reduce(hands, func(a int, h camelHand, i int) int { return a + ((i + 1) * h.bid) }, 0)
	return fmt.Sprintf("%d", res)
}

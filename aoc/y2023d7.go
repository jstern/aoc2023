package aoc

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jstern/aoc2023/aoc/junk"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:7:1", y2023d7part1)
	registerSolution("2023:7:2", y2023d7part2)
}

var camelValues = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}
var jokerValues = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1}

var handStrengths = map[string]int{"5k": 6, "4k": 5, "FH": 4, "3k": 3, "2p": 2, "2k": 1, "HC": 0}

type camelHand struct {
	cards    string
	effcards string
	typ      string
	bid      int
}

func compareHands(h1, h2 camelHand, values map[rune]int) bool {
	// returns true if h2 is higher rank than h1
	h2s := handStrengths[h2.typ]
	h1s := handStrengths[h1.typ]
	if h2s != h1s {
		return h2s > h1s
	}
	h1Vals := lo.Map([]rune(h1.cards), func(r rune, _ int) int { return values[r] })
	h2Vals := lo.Map([]rune(h2.cards), func(r rune, _ int) int { return values[r] })
	for i, h2v := range h2Vals {
		h1v := h1Vals[i]
		if h1v != h2v {
			return h2v > h1v
		}
	}
	return true
}

func handType(cards string) string {
	counter := make(map[rune]int)
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
	panic("can't figure out hand type >" + cards)
}

func bestHandType(cards string) (string, string) {
	nonjokers := strings.ReplaceAll(cards, "J", "")
	jokers := len(cards) - len(nonjokers)
	if jokers == 0 {
		return handType(nonjokers), cards
	}

	best := camelHand{}

	for _, sub := range substitutions("", jokers, lo.Filter(lo.Keys(jokerValues), func(r rune, _ int) bool { return r != 'J' })).Values() {
		effcards := nonjokers + sub
		hand := camelHand{cards: cards, effcards: effcards, typ: handType(effcards)}
		if best.typ == "" {
			best = hand
			continue
		}
		better := compareHands(best, hand, jokerValues)
		if better {
			best = hand
		}
	}
	return best.typ, best.effcards
}

func substitutions(pre string, length int, choices []rune) junk.Set[string] {
	if len(pre) == length {
		return junk.NewSet(pre)
	}
	res := junk.NewSet[string]()
	for _, c := range choices {
		res = res.Union(substitutions(pre+string(c), length, choices))
	}
	return res
}

func parseHand(desc string, jokers bool) camelHand {
	parts := strings.Fields(desc)
	cards := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	ht := handType(cards)
	effcards := cards
	if jokers {
		ht, effcards = bestHandType(cards)
	}
	return camelHand{
		cards:    cards,
		effcards: effcards,
		typ:      ht,
		bid:      bid,
	}
}

func y2023d7part1(input string) string {
	hands := sortedHands(input, false)
	res := lo.Reduce(hands, func(a int, h camelHand, i int) int { return a + ((i + 1) * h.bid) }, 0)
	return fmt.Sprintf("%d", res)
}

func y2023d7part2(input string) string {
	hands := sortedHands(input, true)
	res := lo.Reduce(hands, func(a int, h camelHand, i int) int { return a + ((i + 1) * h.bid) }, 0)
	return fmt.Sprintf("%d", res)
}

func sortedHands(input string, useJokers bool) []camelHand {
	strengths := camelValues
	if useJokers {
		strengths = jokerValues
	}
	hands := lo.Map(strings.Split(strings.TrimSuffix(input, "\n"), "\n"), func(s string, _ int) camelHand {
		return parseHand(s, useJokers)
	})
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], strengths)
	})
	for i, h := range hands {
		if h.cards != h.effcards {
			if h.typ == "HC" || h.typ == "2p" {
				color.Red("%s %s %3d * %4d = %d", h.cards, h.typ, h.bid, i+1, (i+1)*h.bid)
			} else {
				color.Yellow("%s %s %3d * %4d = %d", h.cards, h.typ, h.bid, i+1, (i+1)*h.bid)
			}
		} else {
			fmt.Printf("%s %s %3d * %4d = %d\n", h.cards, h.typ, h.bid, i+1, (i+1)*h.bid)
		}
	}
	return hands
}

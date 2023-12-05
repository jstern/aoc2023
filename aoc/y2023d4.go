package aoc

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:4:1", y2023d4part1)
	registerSolution("2023:4:2", y2023d4part2)
}

func y2023d4part1(input string) string {
	tot := 0.0
	for _, desc := range strings.Split(strings.TrimSpace(input), "\n") {
		hits := scratchcardHits(desc)
		if hits > 0 {
			tot += math.Pow(2, float64(hits)-1)
		}
	}
	return fmt.Sprintf("%d", int(tot))
}

func scratchcardHits(desc string) int {
	_, rest := func() (string, string) { pts := strings.Split(desc, ": "); return pts[0], pts[1] }()
	ws, ps := func() (string, string) { pts := strings.Split(rest, " | "); return pts[0], pts[1] }()
	drawn := junk.NewSet(lo.Map(strings.Fields(ws), func(s string, _ int) int { v, _ := strconv.Atoi(s); return v })...)
	picks := junk.NewSet(lo.Map(strings.Fields(ps), func(s string, _ int) int { v, _ := strconv.Atoi(s); return v })...)
	hits := drawn.Intersection(picks)
	return len(hits)
}

func y2023d4part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cards := make([]*card, len(lines))
	for i, s := range lines {
		cards[i] = &card{id: i + 1, hits: scratchcardHits(s), copies: 0}
	}
	for i := range cards {
		processScratchcard(cards, i)
	}
	tot := lo.SumBy(cards, func(c *card) int { return c.copies })
	return fmt.Sprintf("%d", tot)
}

type card struct {
	id     int
	hits   int
	copies int
}

func processScratchcard(cards []*card, idx int) {
	card := cards[idx]
	card.copies += 1
	for copyIdx := idx + 1; copyIdx < idx+card.hits+1 && copyIdx < len(cards); copyIdx++ {
		processScratchcard(cards, copyIdx)
	}
}

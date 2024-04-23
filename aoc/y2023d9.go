package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:9:1", y2023d9part1)
	registerSolution("2023:9:2", y2023d9part2)
}

func y2023d9part1(input string) string {
	return fmt.Sprint(sumterpolate(input, extrapolate))
}

func sumterpolate(input string, fn func([]int) int) int {
	res := 0
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		seq := lo.Map(strings.Fields(line), func(s string, _ int) int { n, _ := strconv.Atoi(s); return n })
		ex := fn(seq)
		res += ex
		fmt.Println(seq, ex, res)
	}
	return res
}

func reduceseq(seq []int) [][]int {
	seqs := [][]int{seq}
	for {
		ns, done := nextseq(seqs[len(seqs)-1])
		seqs = append(seqs, ns)
		if done {
			return seqs
		}
	}
}

func extrapolate(seq []int) int {
	seqs := reduceseq(seq)
	e := 0
	for _, seq := range lo.Reverse(seqs) {
		e = e + seq[len(seq)-1]
	}
	return e
}

func nextseq(seq []int) (res []int, done bool) {
	res = make([]int, len(seq)-1)
	allzero := true
	for i := 0; i < len(res); i++ {
		d := seq[i+1] - seq[i]
		res[i] = d
		allzero = allzero && d == 0
	}
	return res, allzero
}

func backstrapolate(seq []int) int {
	seqs := reduceseq(seq)
	e := 0
	for _, seq := range lo.Reverse(seqs) {
		e = seq[0] - e
	}
	return e
}

func y2023d9part2(input string) string {
	return fmt.Sprint(sumterpolate(input, backstrapolate))
}

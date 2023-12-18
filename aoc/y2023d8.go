package aoc

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:8:1", y2023d8part1)
	registerSolution("2023:8:2", y2023d8part2)

	nnnRE = regexp.MustCompile(`^(...) = \((...), (...)\)$`)
}

func y2023d8part1(input string) string {
	nn := nnn(strings.Split(strings.TrimSuffix(input, "\n"), "\n"))
	cur := ""
	for {
		if cur == "ZZZ" {
			break
		}
		cur = nn.next()
	}
	return fmt.Sprintf("%d", nn.iteration)
}

var nnnRE *regexp.Regexp

type netscapenode struct {
	key string
	lft string
	rgt string
}

// new netscape navigator node
func nnnn(desc string) netscapenode {
	match := nnnRE.FindStringSubmatch(desc)
	return netscapenode{key: match[1], lft: match[2], rgt: match[3]}
}

func (n netscapenode) String() string {

	res := fmt.Sprintf("%s = (%s, %s)", n.key, n.lft, n.rgt)
	if n.endsIn("A") {
		res = color.GreenString(res)
	}
	if n.endsIn("Z") {
		res = color.RedString(res)
	}
	return res
}

func (n netscapenode) get(ins string) string {
	switch ins {
	case "L":
		return n.lft
	case "R":
		return n.rgt
	}
	panic(ins)
}

func (n netscapenode) endsIn(ltr string) bool {
	return strings.HasSuffix(n.key, ltr)
}

type netscapenavigator struct {
	iteration    int
	instructions []string
	nodes        map[string]netscapenode
	current      netscapenode
}

// new netscape navigator!
func nnn(lines []string) *netscapenavigator {
	nn := netscapenavigator{
		iteration:    0,
		instructions: strings.Split(lines[0], ""),
		nodes:        make(map[string]netscapenode),
	}

	for _, line := range lines[2:] {
		n := nnnn(line)
		if n.key == "AAA" {
			nn.current = n
		}
		nn.nodes[n.key] = n
	}
	return &nn
}

func (nn *netscapenavigator) next() string {
	idx := nn.iteration % len(nn.instructions)
	ins := nn.instructions[idx]
	cur := nn.current
	res := cur.get(ins)
	nn.iteration += 1
	nn.current = nn.nodes[res]
	return res
}

func y2023d8part2(input string) string {
	nn := nnn2(strings.Split(strings.TrimSuffix(input, "\n"), "\n"))
	periods := make([]int, len(nn.current))
	for {
		for i, n := range nn.current {
			if n.endsIn("Z") {
				fmt.Printf("%9d %s\n", nn.iteration, nn.current)
				periods[i] = nn.iteration
			}
		}
		if lo.EveryBy(periods, func(v int) bool { return v != 0 }) {
			break
		}
		nn.next()
	}
	res := lo.Reduce(periods, func(a, b, _ int) int { return lcm(a, b) }, 1)
	return fmt.Sprintf("%d", res)
}

type netscapenavigator2 struct {
	iteration    int
	instructions []string
	nodes        map[string]netscapenode
	current      []netscapenode
}

// new netscape navigator!
func nnn2(lines []string) *netscapenavigator2 {
	nn := netscapenavigator2{
		iteration:    0,
		instructions: strings.Split(lines[0], ""),
		nodes:        make(map[string]netscapenode),
	}

	for _, line := range lines[2:] {
		n := nnnn(line)
		if n.endsIn("A") {
			nn.current = append(nn.current, n)
		}
		nn.nodes[n.key] = n
	}
	return &nn
}

func (nn *netscapenavigator2) next() {
	idx := nn.iteration % len(nn.instructions)
	ins := nn.instructions[idx]
	nn.iteration += 1
	nn.current = lo.Map(nn.current, func(n netscapenode, _ int) netscapenode {
		return nn.nodes[n.get(ins)]
	})
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

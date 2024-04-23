package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:15:1", y2023d15part1)
	registerSolution("2023:15:2", y2023d15part2)
}

func y2023d15part1(input string) string {
	res := lo.Reduce(strings.Split(strings.TrimSuffix(input, "\n"), ","), func(a int, s string, _ int) int { return a + HASH(s) }, 0)
	return fmt.Sprint(res)
}

func HASH(s string) int {
	val := 0
	for _, c := range s {
		val = ((val + int(c)) * 17) % 256
	}
	return val
}

func y2023d15part2(input string) string {
	boxes := make([]lensbox, 256)

	for _, s := range strings.Split(strings.TrimSuffix(input, "\n"), ",") {
		if VerboseEnabled() {
			fmt.Println("-")
			for i, b := range boxes {
				if b.first != nil {
					fmt.Printf("Box %3d:%s\n", i, b)
				}
			}
		}
		parts := strings.Split(s, "=")
		if len(parts) == 2 {
			// it's an add op
			label := parts[0]
			fl, _ := strconv.Atoi(parts[1])
			boxes[HASH(label)].add(label, fl)
			continue
		}
		parts = strings.Split(s, "-")
		label := parts[0]
		boxes[HASH(label)].remove(label)
	}

	if VerboseEnabled() {
		fmt.Println("-")
		for i, b := range boxes {
			if b.first != nil {
				fmt.Printf("Box %3d:%s\n", i, b)
			}
		}
	}

	fp := lo.Reduce(boxes, func(a int, lb lensbox, i int) int { return a + lb.focusingPower(i) }, 0)
	return fmt.Sprint(fp)
}

type lensbox struct {
	first *lens
}

func (lb *lensbox) add(label string, fl int) {
	first := lb.first
	if first == nil {
		lb.first = &lens{label: label, focalLength: fl}
		return
	}
	l := first
	for {
		if l.label == label {
			l.focalLength = fl
			return
		}
		if l.next == nil {
			l.next = &lens{label: label, focalLength: fl}
			return
		}
		l = l.next
	}
}

func (lb *lensbox) remove(label string) {
	first := lb.first
	if first == nil {
		lb.first = nil
		return
	}

	if first.label == label {
		if first.next == nil {
			lb.first = nil
			return
		}
		lb.first = first.next
		return
	}

	prev := first
	l := first.next
	var tgt *lens

	for {
		if l == nil {
			return
		}
		if tgt != nil {
			tgt.next = l
			return
		}
		if l.label == label {
			prev.next = l.next
			return
		}
		prev = l
		l = l.next
	}
}

func (lb lensbox) String() string {
	var b strings.Builder

	l := lb.first

	for {
		if l == nil {
			break
		}
		fmt.Fprintf(&b, " [%s %d]", l.label, l.focalLength)
		l = l.next
	}

	return b.String()
}

func (lb lensbox) lenses() []lens {
	res := make([]lens, 0)

	l := lb.first
	for {
		if l == nil {
			break
		}
		res = append(res, *l)
		l = l.next
	}

	return res
}

func (lb lensbox) focusingPower(box int) int {
	return lo.Reduce(lb.lenses(), func(a int, l lens, i int) int { return a + l.focusingPower(box, i+1) }, 0)
}

type lens struct {
	label       string
	focalLength int
	next        *lens
}

func (l lens) focusingPower(box, pos int) int {
	return (1 + box) * pos * l.focalLength
}

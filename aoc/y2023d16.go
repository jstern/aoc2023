package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:16:1", y2023d16part1)
	registerSolution("2023:16:2", y2023d16part2)
}

func y2023d16part1(input string) string {
	h := newHeater(input)
	h.beams = make(map[int]*beam)
	h.beams[0] = &beam{dx: 1, dy: 0, path: []point{{x: -1, y: 0}}}
	bi := 1
	for {
		beams := lo.ToPairs(h.beams)

		for _, bp := range beams {
			// get next tile for beam. if it's out of bounds, remove it from heater.beams and continue
			tile := bp.Value.next()
			if !h.valid(tile, bp.Value.vec()) {
				delete(h.beams, bp.Key)
				continue
			}

			// otherwise it's in bounds
			newbeam := h.receive(tile, bp.Value)
			if newbeam != nil {
				h.beams[bi] = newbeam
				bi += 1
			}
		}

		fmt.Println(junk.SideBySide("  ", h.PathString(), h.EnergyString()))
		fmt.Println()

		if len(h.beams) == 0 {
			break
		}
	}

	// fmt.Println(junk.SideBySide("  ", h.PathString(), h.EnergyString()))
	// fmt.Println()

	return fmt.Sprintf("%d", h.energizedTiles())
}

type heater struct {
	objects   [][]string
	path      [][]string
	activated [][]bool
	beams     map[int]*beam
	seen      junk.Set[string]
}

func newHeater(input string) heater {
	h := heater{
		objects: lo.Map(strings.Split(strings.TrimSuffix(input, "\n"), "\n"), func(line string, _ int) []string {
			return strings.Split(line, "")
		}),
	}
	h.activated = make([][]bool, len(h.objects))
	for i := 0; i < len(h.activated); i++ {
		h.activated[i] = make([]bool, len(h.objects[0]))
	}
	h.path = make([][]string, len(h.objects))
	for i := 0; i < len(h.activated); i++ {
		h.path[i] = make([]string, len(h.objects[0]))
		copy(h.path[i], h.objects[i])
	}
	h.seen = junk.NewSet[string]()
	return h
}

func (h heater) valid(tile point, vec string) bool {
	// true if the tile is actually on the grid ...
	if tile.y > -1 && tile.y < len(h.objects) && tile.x > -1 && tile.x < len(h.objects) {
		// and we haven't seen a beam hit this tile in the same direction before
		return !h.seen.Contains(vec + tile.String())
	}
	return false
}

func (h heater) receive(tile point, b *beam) *beam {
	// update the path record
	pu := b.vec()
	obj := h.objects[tile.y][tile.x]
	objN, _ := strconv.Atoi(obj)
	if objN != 0 {
		objN += 1
		pu = fmt.Sprintf("%d", objN)
	} else if obj != "." {
		pu = obj
	}
	h.path[tile.y][tile.x] = pu

	// activate the tile
	h.activated[tile.y][tile.x] = true

	h.seen.Add(b.vec() + tile.String())

	// bounce the beam off whatever's in the tile, if anything
	return b.traverse(tile.x, tile.y, obj)
}

func (h heater) energizedTiles() int {
	res := 0
	for _, r := range h.activated {
		res += len(lo.Filter(r, func(v bool, _ int) bool { return v }))
	}

	return res
}

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type beam struct {
	dx   int
	dy   int
	path []point
}

func (b beam) vec() string {
	if b.dx == 1 {
		return ">"
	}
	if b.dx == -1 {
		return "<"
	}
	if b.dy == -1 {
		return "Ʌ"
	}
	if b.dy == 1 {
		return "V"
	}
	return "?"
}

func (b *beam) next() point {
	cur, err := lo.Last(b.path)
	if err != nil {
		panic(err)
	}
	return point{x: cur.x + b.dx, y: cur.y + b.dy}
}

func (b *beam) up() {
	b.dx = 0
	b.dy = -1
}

func (b *beam) down() {
	b.dx = 0
	b.dy = 1
}

func (b *beam) left() {
	b.dx = -1
	b.dy = 0
}

func (b *beam) right() {
	b.dx = 1
	b.dy = 0
}

func (b *beam) traverse(x, y int, obj string) *beam {
	c := b.vec() + obj

	var nb *beam

	// hey dummy! UP = -1

	switch c {
	case ">|", "<|":
		// send this one up and return a new one down
		b.up()
		nb = &beam{}
		nb.down()
	case ">/", "<\\":
		// send this one up
		b.up()
	case ">\\", "</":
		// send this one down
		b.down()
	case "Ʌ/", "V\\":
		// send this one right
		b.right()
	case "Ʌ\\", "V/":
		// send this one left
		b.left()
	case "Ʌ-":
		// send this one right and return a new one left
		b.right()
		nb = &beam{}
		nb.left()
	case "V-":
		// send this one left and return a new one right
		b.left()
		nb = &beam{}
		nb.right()
	}
	b.path = append(b.path, point{x, y})
	if nb != nil {
		nb.path = make([]point, len(b.path))
		copy(nb.path, b.path)
	}
	return nb
}

func (h heater) String() string {
	return strings.Join(lo.Map(h.objects, func(l []string, _ int) string { return strings.Join(l, "") }), "\n")
}

func (h heater) PathString() string {
	return strings.Join(lo.Map(h.path, func(l []string, _ int) string { return strings.Join(l, "") }), "\n")
}

func (h heater) EnergyString() string {
	return strings.Join(
		lo.Map(
			h.activated,
			func(l []bool, _ int) string {
				return strings.Join(
					lo.Map(l, func(c bool, _ int) string {
						if c {
							return "#"
						}
						return "."
					}),
					"",
				)
			},
		),
		"\n",
	)
}

func y2023d16part2(input string) string {
	return "wrong again"
}

package aoc

import (
	"fmt"
	"image/color"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:18:1", y2023d18part1)
	registerSolution("2023:18:2", y2023d18part2)

	digRE = *regexp.MustCompile(`^(.) (\d+) \(#(..)(..)(..)\)$`)
	fillRE = *regexp.MustCompile(`#+\.+#+`)
}

var digRE regexp.Regexp
var fillRE regexp.Regexp

func y2023d18part1(input string) string {
	x := 0
	y := 0
	xmax := x
	xmin := x
	ymax := y
	ymin := y

	ds := make([]dig, 0)

	for _, l := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		m := digRE.FindStringSubmatch(l)
		d := newDig(x, y, m[1], m[2], m[3], m[4], m[5])
		ds = append(ds, d)
		//fmt.Printf("\033[38;2;%d;%d:%d;12m%v\033[0m\n", d.paint.R, d.paint.G, d.paint.G, d)
		x = d.end.x
		y = d.end.y
		xmax = lo.Max([]int{xmax, d.start.x, d.end.x})
		xmin = lo.Min([]int{xmin, d.start.x, d.end.x})
		ymax = lo.Max([]int{ymax, d.start.y, d.end.y})
		ymin = lo.Min([]int{ymin, d.start.y, d.end.y})
	}

	xoffset := -xmin
	rowlen := xmax + 1 - xmin

	var digs digs
	digs = make(map[int][]bool)
	for r := ymin; r <= ymax; r++ {
		digs[r] = make([]bool, rowlen)
	}
	vertices := make(map[string]point)
	for _, d := range ds {
		digs.add(d, xoffset)
		vertices[d.start.String()] = d.start
		vertices[d.end.String()] = d.end
	}

	fmt.Println(vertices)

	digs.print()
	digs.excavate(ymin, ymax)
	digs.print()

	return "niope"
}

type dig struct {
	start point
	end   point
	paint color.RGBA
}

func newDig(x int, y int, dir, len, r, g, b string) dig {
	c := color.RGBA{R: hexToUint8(r), G: hexToUint8(g), B: hexToUint8(b), A: 0xff}
	l, _ := strconv.Atoi(len)
	var end point
	switch dir {
	case "R":
		end = point{x: x + l, y: y}
	case "L":
		end = point{x: x - l, y: y}
	case "D":
		end = point{x: x, y: y + l}
	case "U":
		end = point{x: x, y: y - l}
	}
	return dig{start: point{x, y}, end: end, paint: c}
}

func hexToUint8(s string) uint8 {
	i, _ := strconv.ParseUint(s, 16, 8)
	return uint8(i)
}

type digs map[int][]bool

func (d digs) add(dig dig, xoffset int) {
	sy := dig.start.y
	ey := dig.end.y
	if sy > ey {
		sy, ey = ey, sy
	}
	for y := sy; y <= ey; y++ {
		row := d[y]

		sx, ex := dig.start.x, dig.end.x
		if sx > ex {
			sx, ex = ex, sx
		}
		for x := sx; x <= ex; x++ {
			row[x+xoffset] = true
		}
	}
}

func rowstr(row []bool) string {
	return strings.Join(lo.Map(row, func(v bool, _ int) string {
		if v {
			return "#"
		}
		return "."
	}), "")
}

func (d digs) print() {
	fmt.Println("-")
	keys := lo.Keys(d)
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%4d %s\n", k, rowstr(d[k]))
	}
}

func (d digs) excavate(minY, maxY int) {
	for k := range d {
		if k == minY || k == maxY {
			// leave top and bottom rows alone
			continue
		}
		excavateRow(d[k], rowstr(d[k]))
	}
}

func excavateRow(row []bool, rowstr string) {
	runs := fillRE.FindAllStringIndex(rowstr, -1)
	for _, run := range runs {
		for i := run[0]; i < run[1]; i++ {
			row[i] = true
		}
	}
}

func y2023d18part2(input string) string {
	return "wrong again"
}

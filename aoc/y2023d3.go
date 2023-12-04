package aoc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/jstern/aoc2023/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2023:3:1", y2023d3part1)
	registerSolution("2023:3:2", y2023d3part2)
}

func y2023d3part1(input string) string {
	schematic := parseSchematic(input)
	return fmt.Sprintf("%d", schematic.partsSum)
}

func parseSchematic(desc string) *engineSchematic {
	es := &engineSchematic{
		raw:     strings.TrimSpace(desc),
		symbols: make(map[string]string),
		nums:    make(map[string]schematicNum),
		stars:   make(map[string]junk.Set[string]),
	}

	dr := -1
	ds := -1
	digits := make([]rune, 0)

	finish := func(dr, ds int, digits []rune) (int, int, []rune) {
		if dr >= 0 && ds >= 0 && len(digits) > 0 {
			es.addNumber(dr, ds, digits)
		}
		return -1, -1, make([]rune, 0)
	}

	for row, line := range strings.Split(desc, "\n") {
		for col, chr := range line {
			if unicode.IsDigit(chr) {
				if len(digits) == 0 {
					dr = row
					ds = col
				}
				digits = append(digits, chr)
				continue
			}

			// it wasn't a digit. is it a symbol?
			if chr != '.' {
				es.symbols[schematicKey(row, col)] = string([]rune{chr})
			}

			// end of char: if we got here and we have digits, add them to the
			// list of potential parts in the schematic if we don't already know
			dr, ds, digits = finish(dr, ds, digits)
		}
		// end of line ... if we have a part, finish it out
		dr, ds, digits = finish(dr, ds, digits)
	}
	// end of lines, yadda yadda
	finish(dr, ds, digits)

	es.sumParts()
	es.sumGearRatios()

	if VerboseEnabled() {
		fmt.Println(es)
	}

	return es
}

type engineSchematic struct {
	raw          string
	symbols      map[string]string
	nums         map[string]schematicNum
	stars        map[string]junk.Set[string]
	partsSum     int
	gearRatioSum uint64
}

func (es engineSchematic) String() string {
	str := ""
	for row, line := range strings.Split(es.raw, "\n") {
		hl := 0
		str += fmt.Sprintf("%3d: ", row)
		for col, chr := range strings.Split(line, "") {
			if hl > 0 {
				hl--
				continue
			}
			key := schematicKey(row, col)
			sym := es.symbols[key]
			if sym != "" {
				if es.gearRatio(key) > 0 {
					str += color.HiMagentaString("%s", chr)
				} else {
					str += color.HiYellowString("%s", chr)
				}
				continue
			}

			num := es.nums[key]
			if num.length != 0 {
				hl = num.length - 1
				if es.isPart(num) {
					str += color.HiGreenString("%d", num.value)
				} else {
					str += color.RedString("%d", num.value)
				}
				continue
			}

			str += chr
		}
		str += "\n"
	}
	return str
}

func (es *engineSchematic) addNumber(row, col int, digits []rune) {
	es.nums[schematicKey(row, col)] = newSchematicNum(row, col, digits)
}

func (es *engineSchematic) isPart(sn schematicNum) bool {
	partKey := schematicKey(sn.row, sn.start)
	res := false
	for _, k := range sn.adjacentKeys() {
		sym := es.symbols[k]
		if sym != "" {
			res = true
			if sym == "*" {
				// yes yuck side effects
				parts := lo.ValueOr(es.stars, k, junk.NewSet[string]())
				parts.Add(partKey)
				es.stars[k] = parts
			}
		}
	}
	return res
}

func (es *engineSchematic) sumParts() {
	for _, n := range es.nums {
		if es.isPart(n) {
			es.partsSum += n.value
		}
	}
}

func (es *engineSchematic) sumGearRatios() {
	for k := range es.stars {
		es.gearRatioSum += uint64(es.gearRatio(k))
	}
}

func (es *engineSchematic) gearRatio(key string) int {
	parts := es.stars[key]
	if len(parts) == 2 {
		return lo.Reduce(
			parts.Values(),
			func(a int, k string, _ int) int {
				return a * es.nums[k].value
			},
			1,
		)
	}
	return 0
}

type schematicNum struct {
	row    int
	start  int
	length int
	value  int
}

func newSchematicNum(row, start int, digits []rune) schematicNum {
	val, _ := strconv.Atoi(string(digits))
	return schematicNum{
		row:    row,
		start:  start,
		length: len(digits),
		value:  val,
	}
}

func schematicKey(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
}

func adjacentKeys(row, start, length int) []string {
	res := make([]string, 0)
	prv := row - 1
	nxt := row + 1
	bgn := start - 1
	end := start + length + 1

	for r := prv; r <= nxt; r++ {
		for col := bgn; col < end; col++ {
			// TODO: adding the ones that are *inside* the thing is wasteful,
			// so see if you can exclude them without screwing things up
			// oh but actually, this isn't why your runs take longer ...
			// they take longer because you enabled verbose mode :)
			res = append(res, schematicKey(r, col))
		}
	}
	return res
}

func (sn schematicNum) adjacentKeys() []string {
	return adjacentKeys(sn.row, sn.start, sn.length)
}

func y2023d3part2(input string) string {
	schematic := parseSchematic(input)
	return fmt.Sprintf("%d", schematic.gearRatioSum)
}

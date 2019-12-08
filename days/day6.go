package days

import (
	"regexp"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day6Part1(in string) string {
	reg := regexp.MustCompile(`^((turn on)|(turn off)|(toggle)) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$`)

	lights := map[utils.Coord]bool{}

	for _, instr := range utils.Lines(in) {
		parts := reg.FindStringSubmatch(instr)

		minX := utils.MustAtoi(parts[5])
		minY := utils.MustAtoi(parts[6])
		maxX := utils.MustAtoi(parts[7])
		maxY := utils.MustAtoi(parts[8])

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				switch parts[1] {
				case "turn on":
					lights[utils.Coord{x, y}] = true
				case "turn off":
					lights[utils.Coord{x, y}] = false
				case "toggle":
					current := lights[utils.Coord{x, y}]
					if current {
						lights[utils.Coord{x, y}] = false
					} else {
						lights[utils.Coord{x, y}] = true
					}
				}
			}
		}

	}

	onCount := 0
	for _, v := range lights {
		if v {
			onCount++
		}
	}
	return strconv.Itoa(onCount)
}

func (r *Runner) Day6Part2(in string) string {
	reg := regexp.MustCompile(`^((turn on)|(turn off)|(toggle)) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)$`)

	lights := map[utils.Coord]int{}

	for _, instr := range utils.Lines(in) {
		parts := reg.FindStringSubmatch(instr)

		minX := utils.MustAtoi(parts[5])
		minY := utils.MustAtoi(parts[6])
		maxX := utils.MustAtoi(parts[7])
		maxY := utils.MustAtoi(parts[8])

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {

				switch parts[1] {
				case "turn on":
					lights[utils.Coord{x, y}]++
				case "turn off":
					lights[utils.Coord{x, y}]--
				case "toggle":
					lights[utils.Coord{x, y}] += 2
				}
				current := lights[utils.Coord{x, y}]
				if current < 0 {
					lights[utils.Coord{x, y}] = 0
				}
			}
		}

	}

	brightnessVal := 0
	for _, v := range lights {
		brightnessVal += v
	}
	return strconv.Itoa(brightnessVal)
}

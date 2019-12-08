package days

import "github.com/fivegreenapples/AOC2015/utils"

import "strconv"

import "fmt"

func (r *Runner) Day3Part1(in string) string {
	currentPos := utils.Coord{0, 0}
	houses := map[utils.Coord]int{}
	houses[currentPos] = 1
	for _, instr := range in {
		currentPos = move(currentPos, instr)
		houses[currentPos] = houses[currentPos] + 1
	}
	if r.verbose {
		fmt.Println(houses)
	}
	return strconv.Itoa(len(houses))
}

func (r *Runner) Day3Part2(in string) string {
	currentSantaPos := utils.Coord{0, 0}
	currentRoboPos := utils.Coord{0, 0}
	houses := map[utils.Coord]int{}
	houses[utils.Coord{0, 0}] = 2
	for i, instr := range in {
		if i%2 == 0 {
			currentSantaPos = move(currentSantaPos, instr)
			houses[currentSantaPos] = houses[currentSantaPos] + 1
		} else {
			currentRoboPos = move(currentRoboPos, instr)
			houses[currentRoboPos] = houses[currentRoboPos] + 1
		}

	}
	if r.verbose {
		fmt.Println(houses)
	}
	return strconv.Itoa(len(houses))
}

func move(pos utils.Coord, instruction rune) utils.Coord {
	switch instruction {
	case '<':
		return pos.Add(utils.Coord{-1, 0})
	case '>':
		return pos.Add(utils.Coord{1, 0})
	case '^':
		return pos.Add(utils.Coord{0, 1})
	case 'v':
		return pos.Add(utils.Coord{0, -1})
	}
	return pos
}

package days

import "strconv"

func (r *Runner) Day1Part1(in string) string {
	currentFloor := 0
	for _, instr := range in {
		if instr == '(' {
			currentFloor++
		} else if instr == ')' {
			currentFloor--
		}
	}
	return strconv.Itoa(currentFloor)
}

func (r *Runner) Day1Part2(in string) string {
	currentFloor := 0
	for i, instr := range in {
		if instr == '(' {
			currentFloor++
		} else if instr == ')' {
			currentFloor--
		}
		if currentFloor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return "never reaches basement"
}

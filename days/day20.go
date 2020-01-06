package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day20Part1(in string) string {

	minPresents := utils.MustAtoi(in)

	// Worst case house is if the house is prime so gets presents from elf 1 and elf==house
	// The point is that there will be a house before this that has more presents than this.
	upperBoundHouse := minPresents/10 - 1
	presentsByHouse := make([]int, upperBoundHouse+1)

	for elf := 1; elf <= upperBoundHouse; elf++ {
		for house := elf; house <= upperBoundHouse; house += elf {
			presentsByHouse[house] += (elf * 10)
		}
	}

	for houseNum, presentCount := range presentsByHouse {
		if presentCount > minPresents {
			return strconv.Itoa(houseNum)
		}
	}

	return "Upper bound house not enough"
}

func (r *Runner) Day20Part2(in string) string {

	minPresents := utils.MustAtoi(in)

	// Worst case house is if the house is prime so gets presents from elf 1 and elf==house
	// The point is that there will be a house before this that has more presents than this.
	upperBoundHouse := minPresents/11 - 1
	presentsByHouse := make([]int, upperBoundHouse+1)

	for elf := 1; elf <= upperBoundHouse; elf++ {
		drops := 0
		for house := elf; house <= upperBoundHouse; house += elf {
			presentsByHouse[house] += (elf * 11)
			drops++
			if drops == 50 {
				break
			}
		}
	}

	for houseNum, presentCount := range presentsByHouse {
		if presentCount > minPresents {
			return strconv.Itoa(houseNum)
		}
	}

	return "Upper bound house not enough"
}

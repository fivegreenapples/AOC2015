package days

import (
	"math"
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

const eggnog = 150

func (r *Runner) Day17Part1(in string) string {
	containers := utils.LinesAsInts(in)
	sort.Ints(containers)
	utils.IntSliceReverse(containers)

	combinations := day17NumCombinations(containers, eggnog)
	return strconv.Itoa(combinations)
}

func (r *Runner) Day17Part2(in string) string {
	containers := utils.LinesAsInts(in)
	sort.Ints(containers)
	utils.IntSliceReverse(containers)

	combinations := day17NumCombosForMinContainers(containers, eggnog)
	return strconv.Itoa(combinations)
}

func day17NumCombinations(containers []int, quantity int) int {

	if len(containers) == 0 {
		return 0
	}

	numCombinations := 0
	for cIdx, volume := range containers {
		if volume > quantity {
			continue
		}

		if volume == quantity {
			numCombinations++
			continue
		}

		numCombinations += day17NumCombinations(containers[cIdx+1:], quantity-volume)

	}

	return numCombinations
}

func day17CombinationDetails(containers []int, quantity int) [][]int {

	if len(containers) == 0 {
		return nil
	}

	allCombinations := [][]int{}
	for cIdx, volume := range containers {
		if volume > quantity {
			continue
		}

		if volume == quantity {
			allCombinations = append(allCombinations, []int{volume})
			continue
		}

		theseCombos := day17CombinationDetails(containers[cIdx+1:], quantity-volume)
		for _, thisCombo := range theseCombos {
			allCombinations = append(allCombinations, append([]int{volume}, thisCombo...))
		}

	}

	return allCombinations
}

func day17NumCombosForMinContainers(containers []int, quantity int) int {
	comboDetails := day17CombinationDetails(containers, quantity)

	waysByNumContainers := map[int]int{}
	minNumContainers := math.MaxInt64
	for _, combo := range comboDetails {
		minNumContainers = utils.Min(minNumContainers, len(combo))
		waysByNumContainers[len(combo)]++
	}

	return waysByNumContainers[minNumContainers]
}

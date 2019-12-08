package days

import "github.com/fivegreenapples/AOC2015/utils"

import "strconv"

import "sort"

func (r *Runner) Day2Part1(in string) string {
	presentDetails := utils.MultilineCsvToInts(in, "x")
	paperTotal := 0
	for _, det := range presentDetails {
		sort.Ints(det)
		paperTotal += 3*det[0]*det[1] + 2*det[1]*det[2] + 2*det[2]*det[0]
	}
	return strconv.Itoa(paperTotal)
}

func (r *Runner) Day2Part2(in string) string {
	presentDetails := utils.MultilineCsvToInts(in, "x")
	ribbonTotal := 0
	for _, det := range presentDetails {
		sort.Ints(det)
		ribbonTotal += det[0] + det[0] + det[1] + det[1] + det[0]*det[1]*det[2]
	}
	return strconv.Itoa(ribbonTotal)
}

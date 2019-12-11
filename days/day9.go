package days

import (
	"fmt"
	"math"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day9Part1(in string) string {
	distances := convertDistanceDetailsToDistanceMap(in)

	shortest, _ := calcDistances(distances)
	return fmt.Sprintf("%d", shortest)
}

func (r *Runner) Day9Part2(in string) string {
	distances := convertDistanceDetailsToDistanceMap(in)

	_, longest := calcDistances(distances)
	return fmt.Sprintf("%d", longest)
}

func convertDistanceDetailsToDistanceMap(in string) map[string]map[string]int {
	distanceDetails := utils.MultilineCsvToStrings(in, " ")

	distances := map[string]map[string]int{}

	for _, details := range distanceDetails {

		curMap := distances[details[0]]
		if curMap == nil {
			curMap = map[string]int{}
		}
		curMap[details[2]] = utils.MustAtoi(details[4])
		distances[details[0]] = curMap

		curMap = distances[details[2]]
		if curMap == nil {
			curMap = map[string]int{}
		}
		curMap[details[0]] = utils.MustAtoi(details[4])
		distances[details[2]] = curMap
	}
	return distances
}
func calcDistances(distances map[string]map[string]int) (int, int) {
	cities := []string{}
	for city := range distances {
		cities = append(cities, city)
	}

	if len(cities) <= 1 {
		return 0, 0
	}

	cityPermutations := permutations(cities)

	shortestDist := math.MaxInt64
	longestDist := math.MinInt64
	for _, theseCities := range cityPermutations {
		thisDist := 0
		prevCity := theseCities[0]
		for _, c := range theseCities[1:] {
			thisDist += distances[prevCity][c]
			prevCity = c
		}
		if thisDist < shortestDist {
			shortestDist = thisDist
		}
		if thisDist > longestDist {
			longestDist = thisDist
		}
	}

	return shortestDist, longestDist
}

func permutations(options []string) [][]string {
	if len(options) == 0 {
		panic("can't do permutations of zero elements")
	}
	if len(options) == 1 {
		return [][]string{
			[]string{
				options[0],
			},
		}
	}

	ret := [][]string{}
	for _, opt := range options {

		// copy original options to new slice filtering out current option
		subOptions := []string{}
		for _, subOpt := range options {
			if subOpt != opt {
				subOptions = append(subOptions, subOpt)
			}
		}

		subPerms := permutations(subOptions)
		for _, perm := range subPerms {
			thisPermutation := append([]string{opt}, perm...)
			ret = append(ret, thisPermutation)
		}
	}

	return ret
}

package days

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day24Part1(in string) string {
	packages := utils.LinesAsInts(in)
	packageSum := utils.Sum(packages...)
	sackWeight := packageSum / 3

	for firstGroupSackSize := 1; firstGroupSackSize < len(packages); firstGroupSackSize++ {
		if r.verbose {
			fmt.Println("Trying first group sacksize of", firstGroupSackSize)
		}
		possibleFirstGroups := d24FindGroupsOfNForSum(packages, firstGroupSackSize, sackWeight)
		if r.verbose {
			fmt.Println(" - found", len(possibleFirstGroups))
		}
		if len(possibleFirstGroups) == 0 {
			continue
		}

		// sort possibleFirstGroups by quantum entanglement
		sort.Slice(possibleFirstGroups, func(i int, j int) bool {
			return utils.Product(possibleFirstGroups[i]...) < utils.Product(possibleFirstGroups[j]...)
		})

		// now try each possible first group in turn to see if we can find a second and third group
		for _, firstGroup := range possibleFirstGroups {
			firstGroupMap := map[int]bool{}
			for _, p := range firstGroup {
				firstGroupMap[p] = true
			}

			remainingPackages := []int{}
			for _, p := range packages {
				if !firstGroupMap[p] {
					remainingPackages = append(remainingPackages, p)
				}
			}

			for size := 1; size <= len(remainingPackages)/2; size++ {
				if r.verbose {
					fmt.Println("Trying second group sacksize of", size)
				}

				if d24AreAnyGroupsOfNForSum(remainingPackages, size, sackWeight) {
					// If we have any group2s then we will have group3s made of the
					// remaining packages, so we have found our winning group1
					if r.verbose {
						fmt.Println(" - found at least one")
					}
					return strconv.Itoa(utils.Product(firstGroup...))
				}
				if r.verbose {
					fmt.Println(" - found none")
				}

			}
		}

	}

	return "No solution found :-("
}
func (r *Runner) Day24Part2(in string) string {
	packages := utils.LinesAsInts(in)
	packageSum := utils.Sum(packages...)
	sackWeight := packageSum / 4

	for firstGroupSackSize := 1; firstGroupSackSize < len(packages); firstGroupSackSize++ {
		if r.verbose {
			fmt.Println("Trying first group sacksize of", firstGroupSackSize)
		}
		possibleFirstGroups := d24FindGroupsOfNForSum(packages, firstGroupSackSize, sackWeight)
		if r.verbose {
			fmt.Println(" - found", len(possibleFirstGroups))
		}
		if len(possibleFirstGroups) == 0 {
			continue
		}

		// sort possibleFirstGroups by quantum entanglement
		sort.Slice(possibleFirstGroups, func(i int, j int) bool {
			return utils.Product(possibleFirstGroups[i]...) < utils.Product(possibleFirstGroups[j]...)
		})

		// now try each possible first group in turn to see if there exists a second, third and fourth group
		for _, firstGroup := range possibleFirstGroups {
			firstGroupMap := map[int]bool{}
			for _, p := range firstGroup {
				firstGroupMap[p] = true
			}

			remainingPackages := []int{}
			for _, p := range packages {
				if !firstGroupMap[p] {
					remainingPackages = append(remainingPackages, p)
				}
			}

			for secondGroupSackSize := 1; secondGroupSackSize <= len(remainingPackages)/2; secondGroupSackSize++ {
				if r.verbose {
					fmt.Println("Trying second group sacksize of", secondGroupSackSize)
				}
				possibleSecondGroups := d24FindGroupsOfNForSum(remainingPackages, secondGroupSackSize, sackWeight)
				if r.verbose {
					fmt.Println(" - found", len(possibleSecondGroups))
				}
				if len(possibleSecondGroups) == 0 {
					continue
				}

				// now try each possible second group in turn to see if there exists a third and fourth group
				for _, secondGroup := range possibleSecondGroups {
					secondGroupMap := map[int]bool{}
					for _, p := range secondGroup {
						secondGroupMap[p] = true
					}

					nowRemainingPackages := []int{}
					for _, p := range remainingPackages {
						if !secondGroupMap[p] {
							nowRemainingPackages = append(nowRemainingPackages, p)
						}
					}

					for thirdGroupSackSize := 1; thirdGroupSackSize <= len(nowRemainingPackages)/2; thirdGroupSackSize++ {
						if r.verbose {
							fmt.Println("Trying third group sacksize of", thirdGroupSackSize)
						}

						if d24AreAnyGroupsOfNForSum(nowRemainingPackages, thirdGroupSackSize, sackWeight) {
							// If we have any group3s then we will have group4s made of the
							// remaining packages, so we have found our winning group1
							if r.verbose {
								fmt.Println(" - found at least one")
							}
							return strconv.Itoa(utils.Product(firstGroup...))
						}
						if r.verbose {
							fmt.Println(" - found none")
						}

					}
				}
			}
		}

	}

	return "No solution found :-("
}

func d24FindGroupsOfNForSum(packages []int, n int, sum int) [][]int {

	if n == 0 || sum == 0 {
		return nil
	}

	groups := [][]int{}

	for i, w := range packages {

		if n == 1 {
			if w == sum {
				groups = append(groups, []int{w})
			}
		} else {
			subgroups := d24FindGroupsOfNForSum(packages[i+1:], n-1, sum-w)
			for _, subGrp := range subgroups {
				groups = append(groups, append([]int{w}, subGrp...))
			}
		}

	}

	return groups
}

func d24AreAnyGroupsOfNForSum(packages []int, n int, sum int) bool {

	if n == 0 || sum == 0 {
		return false
	}

	for i, w := range packages {

		if n == 1 {
			if w == sum {
				return true
			}
		} else {
			if d24AreAnyGroupsOfNForSum(packages[i+1:], n-1, sum-w) {
				return true
			}
		}

	}

	return false
}

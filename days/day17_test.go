package days

import "testing"

func TestDay17Part1(t *testing.T) {

	numCombos := day17NumCombinations([]int{20, 15, 10, 5, 5}, 25)
	if numCombos != 4 {
		t.Errorf("day17 part 1 failed. Expected 4 combinations from containers of 20,15,10,5,5. Got %d", numCombos)
	}

}

func TestDay17Part2(t *testing.T) {

	numCombos := day17NumCombosForMinContainers([]int{20, 15, 10, 5, 5}, 25)
	if numCombos != 3 {
		t.Errorf("day17 part 2 failed. Expected 3 combinations from containers of 20,15,10,5,5. Got %d", numCombos)
	}

}

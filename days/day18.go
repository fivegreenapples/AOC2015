package days

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day18Part1(in string) string {
	grid := day18ProcessInputToGrid(in)

	if r.verbose {
		day18RenderGrid(grid)
	}
	for i := 1; i <= 100; i++ {
		grid = day18AnimateGrid(grid, false)
		if r.verbose {
			time.Sleep(100 * time.Millisecond)
			day18RenderGrid(grid)
		}
	}

	return strconv.Itoa(day18CountLightsOn(grid))
}

func (r *Runner) Day18Part2(in string) string {
	grid := day18ProcessInputToGrid(in)
	// force corners to ON
	grid[utils.Coord{0, 0}] = true
	grid[utils.Coord{99, 0}] = true
	grid[utils.Coord{0, 99}] = true
	grid[utils.Coord{99, 99}] = true

	if r.verbose {
		day18RenderGrid(grid)
	}
	for i := 1; i <= 100; i++ {
		grid = day18AnimateGrid(grid, true)
		if r.verbose {
			time.Sleep(100 * time.Millisecond)
			day18RenderGrid(grid)
		}
	}

	return strconv.Itoa(day18CountLightsOn(grid))
}

func day18ProcessInputToGrid(in string) map[utils.Coord]bool {
	grid := map[utils.Coord]bool{}
	lines := utils.Lines(in)
	for y, line := range lines {
		for x, char := range line {
			grid[utils.Coord{x, y}] = char == '#'
		}
	}
	return grid
}

func day18RenderGrid(grid map[utils.Coord]bool) {
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			if grid[utils.Coord{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func day18AnimateGrid(grid map[utils.Coord]bool, gameOfLifeMode bool) map[utils.Coord]bool {

	newGrid := map[utils.Coord]bool{}
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {

			thisPos := utils.Coord{x, y}
			current := grid[thisPos]

			neighbour1 := grid[utils.Coord{x - 1, y - 1}]
			neighbour2 := grid[utils.Coord{x, y - 1}]
			neighbour3 := grid[utils.Coord{x + 1, y - 1}]
			neighbour4 := grid[utils.Coord{x - 1, y}]
			neighbour5 := grid[utils.Coord{x + 1, y}]
			neighbour6 := grid[utils.Coord{x - 1, y + 1}]
			neighbour7 := grid[utils.Coord{x, y + 1}]
			neighbour8 := grid[utils.Coord{x + 1, y + 1}]

			numNeighboursOn := 0
			if neighbour1 {
				numNeighboursOn++
			}
			if neighbour2 {
				numNeighboursOn++
			}
			if neighbour3 {
				numNeighboursOn++
			}
			if neighbour4 {
				numNeighboursOn++
			}
			if neighbour5 {
				numNeighboursOn++
			}
			if neighbour6 {
				numNeighboursOn++
			}
			if neighbour7 {
				numNeighboursOn++
			}
			if neighbour8 {
				numNeighboursOn++
			}

			newState := false
			if current {
				newState = numNeighboursOn == 2 || numNeighboursOn == 3
			} else {
				newState = numNeighboursOn == 3
			}

			newGrid[thisPos] = newState
		}
	}

	if gameOfLifeMode {
		// force corners to ON
		newGrid[utils.Coord{0, 0}] = true
		newGrid[utils.Coord{99, 0}] = true
		newGrid[utils.Coord{0, 99}] = true
		newGrid[utils.Coord{99, 99}] = true
	}

	return newGrid
}

func day18CountLightsOn(grid map[utils.Coord]bool) int {
	onCount := 0
	for _, val := range grid {
		if val {
			onCount++
		}
	}
	return onCount
}

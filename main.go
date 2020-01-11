package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/fivegreenapples/AOC2015/days"
)

func main() {

	day := flag.Int("d", 0, "Day of Advent")
	verbose := flag.Bool("v", false, "Verbosity")
	part := flag.Int("p", 0, "Part")
	flag.Parse()

	runner := days.NewRunner(*verbose)

	startDay := 1
	endDay := 25
	if *day > 0 {
		startDay = *day
		endDay = *day
	}
	for d := startDay; d <= endDay; d++ {

		inputFileName := "inputs/day" + strconv.Itoa(d) + ".txt"
		puzzleInputBytes, err := ioutil.ReadFile(inputFileName)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Error: couldn't read input file: %v\n", err)
			os.Exit(2)
		}
		puzzleInput := string(puzzleInputBytes)

		if *part == 0 || *part == 1 {
			pt1Out, pt1err := runner.Run(d, 1, puzzleInput)
			if pt1err != nil {
				fmt.Printf("Error: %s\n", pt1err.Error())
				os.Exit(3)
			}
			fmt.Printf("Day %d part %d: %s\n", d, 1, pt1Out)
		}
		if *part == 0 || *part == 2 {
			pt2Out, pt2err := runner.Run(d, 2, puzzleInput)
			if pt2err == nil {
				fmt.Printf("Day %d part %d: %s\n", d, 2, pt2Out)
			} else if *part == 2 {
				// only display error if we asked for part 2
				fmt.Printf("Error: %s\n", pt2err.Error())
				os.Exit(3)
			}
		}
	}
}

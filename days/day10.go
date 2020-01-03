package days

import (
	"fmt"
	"strconv"
	"strings"
)

func (r *Runner) Day10Part1(in string) string {
	for i := 0; i < 40; i++ {
		if r.verbose {
			fmt.Println(in)
		}
		in = day10Iterate(in)
	}
	if r.verbose {
		fmt.Println(in)
	}
	return strconv.Itoa(len(in))
}

func (r *Runner) Day10Part2(in string) string {
	for i := 0; i < 50; i++ {
		if r.verbose {
			fmt.Println(in)
		}
		in = day10Iterate(in)
	}
	if r.verbose {
		fmt.Println(in)
	}
	return strconv.Itoa(len(in))
}

func day10Iterate(in string) string {
	var ret strings.Builder
	var current rune
	var count int
	for _, rn := range in {
		if rn != current {
			if count > 0 {
				fmt.Fprintf(&ret, "%d%s", count, string(current))
			}
			count = 1
			current = rn
		} else {
			count++
		}
	}
	fmt.Fprintf(&ret, "%d%s", count, string(current))
	return ret.String()
}

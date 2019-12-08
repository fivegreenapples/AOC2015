package days

import "github.com/fivegreenapples/AOC2015/utils"

import "strconv"

import "fmt"

func (r *Runner) Day5Part1(in string) string {
	niceCount := 0
	for _, test := range utils.Lines(in) {
		if r.d5IsNicePt1(test) {
			niceCount++
		}
	}
	return strconv.Itoa(niceCount)
}

func (r *Runner) Day5Part2(in string) string {
	niceCount := 0
	for _, test := range utils.Lines(in) {
		if r.d5IsNicePt2(test) {
			niceCount++
		}
	}
	return strconv.Itoa(niceCount)
}

func (r *Runner) d5IsNicePt1(in string) bool {

	vowelCount := 0
	var previousRune rune
	var vowelTest, consecutiveRuneTest, naughtyPairTest bool
	for _, r := range in {

		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowelCount++
		}

		if r == previousRune {
			consecutiveRuneTest = true
		}

		if previousRune == 'a' && r == 'b' || previousRune == 'c' && r == 'd' || previousRune == 'p' && r == 'q' || previousRune == 'x' && r == 'y' {
			naughtyPairTest = true
		}

		previousRune = r
	}

	vowelTest = vowelCount >= 3

	if r.verbose {
		fmt.Println(in, vowelTest, consecutiveRuneTest, !naughtyPairTest)
	}

	return vowelTest && consecutiveRuneTest && !naughtyPairTest
}

func (r *Runner) d5IsNicePt2(in string) bool {

	var previousRune, previousPreviousRune rune
	var repeatedPairTest, separatedRepeatTest bool
	pairs := map[string]int{}
	for i, r := range in {

		if r == previousPreviousRune {
			separatedRepeatTest = true
		}

		thisPair := string(previousRune) + string(r)
		if idx, seen := pairs[thisPair]; seen {
			if idx < i-2 {
				repeatedPairTest = true
			}
		} else {
			pairs[thisPair] = i - 1
		}

		previousPreviousRune = previousRune
		previousRune = r
	}

	if r.verbose {
		fmt.Println(in, repeatedPairTest, separatedRepeatTest)
	}

	return repeatedPairTest && separatedRepeatTest
}

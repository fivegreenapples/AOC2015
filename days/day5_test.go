package days

import "testing"

func TestDay5Part1(t *testing.T) {

	testInputs := map[string]bool{
		"ugknbfddgicrmopn": true,  // because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
		"aaa":              true,  // because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
		"jchzalrnumimnmhp": false, // because it has no double letter.
		"haegwjzuvuyypxyu": false, // because it contains the string xy.
		"dvszwmarrgswjxmb": false, // because it contains only one vowel.
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out := dayRunner.d5IsNicePt1(in)
		if out != expectedOut {
			t.Errorf("day5 isNicePt1 test failed with %s. Expected %v, got %v", in, expectedOut, out)
		}
	}

}

func TestDay5Part2(t *testing.T) {

	testInputs := map[string]bool{
		"qjhvhtzxzqqjkmpb": true,  // because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
		"xxyxx":            true,  // because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
		"uurcxstgmygtbstg": false, // because it has a pair (tg) but no repeat with a single letter between them.
		"ieodomkazucvgmuy": false, // because it has a repeating letter with one between (odo), but no pair that appears twice.
		"aaa":              false, // because it has a repeating letter with one between (aaa), but no pair that appears twice without overlapping.
		"aaaa":             true,  // because it has a repeating letter with one between (aaa), and a pair that appears twice without overlapping.
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out := dayRunner.d5IsNicePt2(in)
		if out != expectedOut {
			t.Errorf("day5 isNicePt2 test failed with %s. Expected %v, got %v", in, expectedOut, out)
		}
	}

}

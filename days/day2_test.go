package days

import "testing"

func TestDay2Part1(t *testing.T) {

	testInputs := map[string]string{
		"2x3x4":  "58",
		"1x1x10": "43",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(2, 1, in)
		if out != expectedOut {
			t.Errorf("day2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay2Part2(t *testing.T) {

	testInputs := map[string]string{
		"2x3x4":  "34",
		"1x1x10": "14",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(2, 2, in)
		if out != expectedOut {
			t.Errorf("day2 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

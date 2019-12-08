package days

import "testing"

func TestDay3Part1(t *testing.T) {

	testInputs := map[string]string{
		">":          "2",
		"^>v<":       "4",
		"^v^v^v^v^v": "2",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 1, in)
		if out != expectedOut {
			t.Errorf("day3 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay3Part2(t *testing.T) {

	testInputs := map[string]string{
		"^v":         "3",
		"^>v<":       "3",
		"^v^v^v^v^v": "11",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 2, in)
		if out != expectedOut {
			t.Errorf("day3 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

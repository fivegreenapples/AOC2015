package days

import "testing"

func TestDay4Part1(t *testing.T) {

	testInputs := map[string]string{
		"abcdef":  "609043",
		"pqrstuv": "1048970",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(4, 1, in)
		if out != expectedOut {
			t.Errorf("day4 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

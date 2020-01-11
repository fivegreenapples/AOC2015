package days

import "testing"

func TestDay24Part1(t *testing.T) {

	testInputs := map[string]string{
		`1
		2
		3
		4
		5
		7
		8
		9
		10
		11`: "99",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(24, 1, in)
		if out != expectedOut {
			t.Errorf("day24 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay24Part2(t *testing.T) {

	testInputs := map[string]string{
		`1
		2
		3
		4
		5
		7
		8
		9
		10
		11`: "44",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(24, 2, in)
		if out != expectedOut {
			t.Errorf("day24 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

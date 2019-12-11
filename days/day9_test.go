package days

import "testing"

func TestDay9Part1(t *testing.T) {

	testInputs := map[string]string{
		`London to Dublin = 464
		London to Belfast = 518
		Dublin to Belfast = 141`: "605",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(9, 1, in)
		if out != expectedOut {
			t.Errorf("day9 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay9Part2(t *testing.T) {

	testInputs := map[string]string{
		`London to Dublin = 464
		London to Belfast = 518
		Dublin to Belfast = 141`: "982",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(9, 2, in)
		if out != expectedOut {
			t.Errorf("day9 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

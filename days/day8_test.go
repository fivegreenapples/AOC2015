package days

import "testing"

func TestDay8Part1(t *testing.T) {

	testInputs := map[string]string{
		`""`:         "2",
		`"abc"`:      "2",
		`"aaa\"aaa"`: "3",
		`"\x27"`:     "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(8, 1, in)
		if out != expectedOut {
			t.Errorf("day8 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay8Part2(t *testing.T) {

	testInputs := map[string]string{
		`""
		"abc"
		"aaa\"aaa"
		"\x27"`: "19",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(8, 2, in)
		if out != expectedOut {
			t.Errorf("day8 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

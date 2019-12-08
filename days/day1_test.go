package days

import "testing"

func TestDay1Part1(t *testing.T) {

	testInputs := map[string]string{
		"(())":    "0",
		"()()":    "0",
		"(((":     "3",
		"(()(()(": "3",
		"))(((((": "3",
		"())":     "-1",
		"))(":     "-1",
		")))":     "-3",
		")())())": "-3",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 1, in)
		if out != expectedOut {
			t.Errorf("day1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay1Part2(t *testing.T) {

	testInputs := map[string]string{
		")":     "1",
		"()())": "5",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(1, 2, in)
		if out != expectedOut {
			t.Errorf("day1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

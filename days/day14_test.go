package days

import "testing"

func TestDay14Part1(t *testing.T) {

	testInputs := map[string]string{
		`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
		Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`: "2660",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(14, 1, in)
		if out != expectedOut {
			t.Errorf("day14 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay14Part2(t *testing.T) {

	testInputs := map[string]string{
		`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
		Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`: "1564",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(14, 2, in)
		if out != expectedOut {
			t.Errorf("day14 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

package days

import "testing"

func TestDay15Part1(t *testing.T) {

	testInputs := map[string]string{
		`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`: "62842880",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(15, 1, in)
		if out != expectedOut {
			t.Errorf("day15 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay15Part2(t *testing.T) {

	testInputs := map[string]string{
		`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`: "57600000",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(15, 2, in)
		if out != expectedOut {
			t.Errorf("day15 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

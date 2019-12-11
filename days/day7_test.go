package days

import "testing"

func TestDay7Part1(t *testing.T) {

	thisCircuit := circuit{
		resolvedWires: map[string]uint16{},
		rawWires: map[string]string{
			"x": "123",
			"y": "456",
			"d": "x AND y",
			"e": "x OR y",
			"f": "x LSHIFT 2",
			"g": "y RSHIFT 2",
			"h": "NOT x",
			"i": "NOT y",
		},
	}

	expected := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	for wire, expectedVal := range expected {
		val := thisCircuit.resolve(wire)

		if val != expectedVal {
			t.Errorf("day7 failed with wire %s. Expected value %d, got %d", wire, expectedVal, val)
		}

	}

}

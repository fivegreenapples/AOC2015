package days

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day7Part1(in string) string {
	wires := utils.MultilineCsvToStrings(in, " -> ")

	circuit := circuit{
		rawWires:      map[string]string{},
		resolvedWires: map[string]uint16{},
	}
	for _, w := range wires {
		circuit.rawWires[w[1]] = w[0]
	}

	valA := circuit.resolve("a")
	return strconv.FormatUint(uint64(valA), 10)
}

func (r *Runner) Day7Part2(in string) string {
	wires := utils.MultilineCsvToStrings(in, " -> ")

	circuit := circuit{
		rawWires:      map[string]string{},
		resolvedWires: map[string]uint16{},
	}
	for _, w := range wires {
		circuit.rawWires[w[1]] = w[0]
	}
	circuit.rawWires["b"] = "46065"

	valA := circuit.resolve("a")
	return strconv.FormatUint(uint64(valA), 10)
}

type circuit struct {
	rawWires      map[string]string
	resolvedWires map[string]uint16
	reg           *regexp.Regexp
}

func (c *circuit) resolve(wire string) uint16 {

	if val, resolved := c.resolvedWires[wire]; resolved {
		return val
	}

	// check if wire is a number
	wireVal, ok := strconv.Atoi(wire)
	if ok == nil {
		c.resolvedWires[wire] = uint16(wireVal)
		return uint16(wireVal)
	}

	details := c.rawWires[wire]

	if c.reg == nil {
		c.reg = regexp.MustCompile(`^(([0-9a-z]+)|(([0-9a-z]+) (AND|OR) ([0-9a-z]+))|(([0-9a-z]+) (LSHIFT|RSHIFT) ([0-9a-z]+))|(NOT ([0-9a-z]+)))$`)
	}

	matches := c.reg.FindStringSubmatch(details)
	if len(matches) == 0 {
		panic(fmt.Errorf("unexpected no regex match for %s", details))
	}

	var val uint16
	if len(matches[2]) > 0 {
		val = c.resolve(matches[2])
	} else if matches[5] == "AND" {
		val = c.resolve(matches[4]) & c.resolve(matches[6])
	} else if matches[5] == "OR" {
		val = c.resolve(matches[4]) | c.resolve(matches[6])
	} else if matches[9] == "LSHIFT" {
		val = c.resolve(matches[8]) << utils.MustAtoi(matches[10])
	} else if matches[9] == "RSHIFT" {
		val = c.resolve(matches[8]) >> utils.MustAtoi(matches[10])
	} else if len(matches[12]) > 0 {
		val = ^c.resolve(matches[12])
	}

	c.resolvedWires[wire] = val
	return val
}

package days

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func (r *Runner) Day12Part1(in string) string {

	buf := bytes.NewBufferString(in)
	d := json.NewDecoder(buf)
	d.UseNumber()

	sumNumber := 0
	tok, err := d.Token()
	for err == nil {
		if number, ok := tok.(json.Number); ok {
			if intNum, err := number.Int64(); err == nil {
				sumNumber += int(intNum)
			}
		}

		tok, err = d.Token()
	}

	return strconv.Itoa(sumNumber)
}

func (r *Runner) Day12Part2(in string) string {
	buf := bytes.NewBufferString(in)
	d := json.NewDecoder(buf)
	d.UseNumber()

	tok, err := d.Token()
	sumsByDepth := map[int]int{}
	ignoresByDepth := map[int]bool{}
	curDepth := 0
	stack := []string{}
	for err == nil {

		switch tokVal := tok.(type) {
		case string:
			if tokVal == "red" {
				if stack[len(stack)-1] == "object" {
					ignoresByDepth[curDepth] = true
				}
			}
		case json.Delim:
			if tokVal.String() == "{" {
				curDepth++
				stack = append(stack, "object")
			} else if tokVal.String() == "}" {
				if !ignoresByDepth[curDepth] {
					sumsByDepth[curDepth-1] += sumsByDepth[curDepth]
				}
				sumsByDepth[curDepth] = 0
				ignoresByDepth[curDepth] = false

				curDepth--
				stack = stack[:len(stack)-1]
			} else if tokVal.String() == "[" {
				stack = append(stack, "array")
			} else if tokVal.String() == "]" {
				stack = stack[:len(stack)-1]
			}
		case json.Number:
			if intNum, err := tokVal.Int64(); err == nil {
				sumsByDepth[curDepth] = sumsByDepth[curDepth] + int(intNum)
			}
		}

		tok, err = d.Token()
	}

	return strconv.Itoa(sumsByDepth[0])
}

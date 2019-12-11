package days

import (
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day8Part1(in string) string {
	inputs := utils.Lines(in)

	totalCode, totalString := 0, 0
	for _, str := range inputs {
		code, str := calcCodeAndStringLength(str)
		totalCode += code
		totalString += str
	}

	return strconv.Itoa(totalCode - totalString)
}

func (r *Runner) Day8Part2(in string) string {
	inputs := utils.Lines(in)
	encodedInputs := []string{}
	for _, input := range inputs {
		encodedInputs = append(encodedInputs, encodeString(input))
	}
	return r.Day8Part1(strings.Join(encodedInputs, "\n"))
}

func calcCodeAndStringLength(in string) (code, str int) {

	code = len(in)
	str = 0

	i := 1
	for i < len(in)-1 {
		char := in[i]
		nextChar := in[i+1]

		switch char {
		case '\\':
			if nextChar == 'x' {
				i += 4
			} else {
				i += 2
			}
		default:
			i++
		}
		str++

	}

	return code, str
}

func encodeString(in string) string {
	out := `"`
	for _, char := range in {
		switch char {
		case '\\', '"':
			out += "\\"
		}
		out += string(char)
	}
	return out + `"`
}

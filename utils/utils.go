package utils

import (
	"math"
	"strconv"
	"strings"
)

func Lines(in string) []string {
	lines := strings.Split(in, "\n")
	trimmed := []string{}
	for _, l := range lines {
		trimmed = append(trimmed, strings.TrimSpace(l))
	}
	return trimmed
}

func LinesAsInts(in string) []int {
	return StringsToInts(Lines(in))
}

func CsvToInts(in string) []int {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, ",")
	bits := strings.Split(in, ",")
	return StringsToInts(bits)
}
func CsvToStrings(in string) []string {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, ",")
	bits := strings.Split(in, ",")
	return bits
}
func MultilineCsvToInts(in string, sep string) [][]int {
	ret := [][]int{}
	for _, l := range Lines(in) {
		l = strings.TrimSpace(l)
		l = strings.Trim(l, sep)
		bits := strings.Split(l, sep)
		ret = append(ret, StringsToInts(bits))
	}
	return ret
}
func MultilineCsvToStrings(in string, sep string) [][]string {
	ret := [][]string{}
	for _, l := range Lines(in) {
		l = strings.TrimSpace(l)
		l = strings.Trim(l, sep)
		bits := strings.Split(l, sep)
		ret = append(ret, bits)
	}
	return ret
}

func StringsToInts(inStrings []string) []int {
	ints := []int{}
	for _, in := range inStrings {
		in := strings.TrimSpace(in)
		thisInt, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		ints = append(ints, thisInt)
	}
	return ints
}

func StringToDigits(in string) []int {
	digits := make([]int, len(in))
	for i, d := range strings.Split(in, "") {
		digits[i] = MustAtoi(d)
	}
	return digits
}

func DigitsToString(in []int) string {
	stringDigits := make([]string, len(in))
	for i, d := range in {
		stringDigits[i] = strconv.Itoa(d)
	}
	return strings.Join(stringDigits, "")
}

func MustAtoi(in string) int {
	ret, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return ret
}

func StringSliceReverse(in []string) []string {
	for left, right := 0, len(in)-1; left < right; left, right = left+1, right-1 {
		in[left], in[right] = in[right], in[left]
	}
	return in
}

func IntSliceReverse(in []int) []int {
	for left, right := 0, len(in)-1; left < right; left, right = left+1, right-1 {
		in[left], in[right] = in[right], in[left]
	}
	return in
}
func Min(in ...int) int {
	minimum := math.MaxInt64
	for _, v := range in {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}
func Max(in ...int) int {
	maximum := math.MinInt64
	for _, v := range in {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

func Sum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}
func Product(vals ...int) int {
	sum := 1
	for _, v := range vals {
		sum *= v
	}
	return sum
}

package days

import "strconv"

func (r *Runner) Day25Part1(in string) string {
	// input: To continue, please consult the code grid in the manual.  Enter the code at row 2947, column 3029.

	row := 2947
	col := 3029

	// Following worked out by inspection
	// row X column N is (N+X-1)th triangle number minus (X-1)
	// triangle number is sum of 1 to N: N(N+1)/2
	triangleNum := col + row - 1
	codeIdx := ((triangleNum * (triangleNum + 1)) / 2) - (row - 1)

	code := 20151125
	for i := 2; i <= codeIdx; i++ {
		code = (code * 252533) % 33554393
	}

	return strconv.Itoa(code)
}

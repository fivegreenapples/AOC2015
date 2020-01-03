package days

import (
	"github.com/fivegreenapples/AOC2015/password"
)

func (r *Runner) Day11Part1(in string) string {

	p := password.Password(in)
	p.IncrementUntilValid()

	return string(p)
}

func (r *Runner) Day11Part2(in string) string {
	p := password.Password(in)
	p.IncrementUntilValid()
	p.IncrementUntilValid()

	return string(p)
}

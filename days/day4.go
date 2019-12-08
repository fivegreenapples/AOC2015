package days

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func (r *Runner) Day4Part1(in string) string {
	hasher := md5.New()
	var hash []byte
	var i int
	for i = 1; i >= 1; i++ {
		fmt.Fprintf(hasher, "%s%d", in, i)
		hash = hasher.Sum(nil)
		if hash[0] == 0 && hash[1] == 0 && hash[2]&0xf0 == 0 {
			break
		}
		hasher.Reset()
	}
	return strconv.Itoa(i)
}

func (r *Runner) Day4Part2(in string) string {
	hasher := md5.New()
	var hash []byte
	var i int
	for i = 1; i >= 1; i++ {
		fmt.Fprintf(hasher, "%s%d", in, i)
		hash = hasher.Sum(nil)
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			break
		}
		hasher.Reset()
	}
	return strconv.Itoa(i)
}

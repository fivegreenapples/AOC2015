package utils

func PermutationsInt(options []int) [][]int {
	if len(options) == 0 {
		panic("can't do permutations of zero elements")
	}
	if len(options) == 1 {
		return [][]int{
			[]int{
				options[0],
			},
		}
	}

	ret := [][]int{}
	for idx, opt := range options {

		// copy original options to new slice filtering out current option
		subOptions := []int{}
		for subIdx, subOpt := range options {
			if idx != subIdx {
				subOptions = append(subOptions, subOpt)
			}
		}

		subPerms := PermutationsInt(subOptions)
		for _, perm := range subPerms {
			thisPermutation := append([]int{opt}, perm...)
			ret = append(ret, thisPermutation)
		}
	}

	return ret
}
func PermutationsString(options []string) [][]string {
	if len(options) == 0 {
		panic("can't do permutations of zero elements")
	}
	if len(options) == 1 {
		return [][]string{
			[]string{
				options[0],
			},
		}
	}

	ret := [][]string{}
	for idx, opt := range options {

		// copy original options to new slice filtering out current option
		subOptions := []string{}
		for subIdx, subOpt := range options {
			if idx != subIdx {
				subOptions = append(subOptions, subOpt)
			}
		}

		subPerms := PermutationsString(subOptions)
		for _, perm := range subPerms {
			thisPermutation := append([]string{opt}, perm...)
			ret = append(ret, thisPermutation)
		}
	}

	return ret
}

package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day19Part1(in string) string {

	details := utils.Lines(in)

	// expect last line to be medicine molecule:
	medicineMolecule := details[len(details)-1]

	// expect other lines to be replacements
	replacementDetails := strings.Join(details[:len(details)-2], "\n")
	replacements := utils.StringsFromRegex(replacementDetails, `^([a-zA-Z]+) => ([a-zA-Z]+)$`)
	replacementMap := map[string][]string{}
	for _, repl := range replacements {
		replacementMap[repl[1]] = append(replacementMap[repl[1]], repl[2])
	}

	// split medicine molecule into atoms
	medicineMoleculeSplit := []string{}
	currentMolecule := string(medicineMolecule[0])
	for _, char := range medicineMolecule[1:] {
		if char >= 'A' && char <= 'Z' {
			medicineMoleculeSplit = append(medicineMoleculeSplit, currentMolecule)
			currentMolecule = string(char)
		} else {
			currentMolecule = currentMolecule + string(char)
		}

	}
	medicineMoleculeSplit = append(medicineMoleculeSplit, currentMolecule)

	// Do one replacement round
	newMolecules := map[string]int{}
	for idx, atom := range medicineMoleculeSplit {

		if replacementMap[atom] == nil {
			continue
		}

		for _, replacement := range replacementMap[atom] {
			thisMolecule := strings.Join(medicineMoleculeSplit[:idx], "")
			thisMolecule = thisMolecule + replacement
			thisMolecule = thisMolecule + strings.Join(medicineMoleculeSplit[idx+1:], "")

			newMolecules[thisMolecule]++
		}

	}

	return strconv.Itoa(len(newMolecules))
}

func (r *Runner) Day19Part2(in string) string {
	details := utils.Lines(in)

	// expect last line to be medicine molecule:
	medicineMolecule := details[len(details)-1]

	// expect other lines to be replacements
	replacementDetails := strings.Join(details[:len(details)-2], "\n")
	replacements := utils.StringsFromRegex(replacementDetails, `^([a-zA-Z]+) => ([a-zA-Z]+)$`)
	replacementReverseMap := map[string]string{}
	replacementList := []string{}
	for _, repl := range replacements {
		replacementReverseMap[repl[2]] = repl[1]
		replacementList = append(replacementList, repl[2])
	}

	// sort replacements by length descending
	sort.Slice(replacementList, func(i, j int) bool {
		return len(replacementList[i]) > len(replacementList[j])
	})

	// Loop over sorted replacements and replace if we can, always trying the longest
	// replacement string. We are working backwards from the medicine molecule to the
	// bare electron. Stop if we get to a single "e".
	//
	// It's not clear to me that this guarantees getting to an "e" let alone it being the
	// shortest number of steps. But for my input at least, this seems to yield the correct
	// answer.
	steps := 0
	for {
		if r.verbose {
			fmt.Println(steps, "-", medicineMolecule)
		}
		if medicineMolecule == "e" {
			break
		}
		for _, repl := range replacementList {
			if strings.Index(medicineMolecule, repl) == -1 {
				continue
			}

			medicineMolecule = strings.Replace(medicineMolecule, repl, replacementReverseMap[repl], 1)
			steps++
			break
		}
	}

	return strconv.Itoa(steps)
}

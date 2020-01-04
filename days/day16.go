package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

var giftDetail = sue{
	children:    3,
	cats:        7,
	samoyeds:    2,
	pomeranians: 3,
	akitas:      0,
	vizslas:     0,
	goldfish:    5,
	trees:       3,
	cars:        2,
	perfumes:    1,
}

func (r *Runner) Day16Part1(in string) string {
	sues := day16ExtractSues(in)
	for sueIdx, thisSue := range sues {

		if thisSue.children >= 0 && thisSue.children != giftDetail.children {
			continue
		}
		if thisSue.cats >= 0 && thisSue.cats != giftDetail.cats {
			continue
		}
		if thisSue.samoyeds >= 0 && thisSue.samoyeds != giftDetail.samoyeds {
			continue
		}
		if thisSue.pomeranians >= 0 && thisSue.pomeranians != giftDetail.pomeranians {
			continue
		}
		if thisSue.akitas >= 0 && thisSue.akitas != giftDetail.akitas {
			continue
		}
		if thisSue.vizslas >= 0 && thisSue.vizslas != giftDetail.vizslas {
			continue
		}
		if thisSue.goldfish >= 0 && thisSue.goldfish != giftDetail.goldfish {
			continue
		}
		if thisSue.trees >= 0 && thisSue.trees != giftDetail.trees {
			continue
		}
		if thisSue.cars >= 0 && thisSue.cars != giftDetail.cars {
			continue
		}
		if thisSue.perfumes >= 0 && thisSue.perfumes != giftDetail.perfumes {
			continue
		}

		return strconv.Itoa(sueIdx + 1)
	}
	return "no sue found"
}

func (r *Runner) Day16Part2(in string) string {
	sues := day16ExtractSues(in)
	for sueIdx, thisSue := range sues {

		if thisSue.children >= 0 && thisSue.children != giftDetail.children {
			continue
		}
		if thisSue.cats >= 0 && thisSue.cats <= giftDetail.cats {
			continue
		}
		if thisSue.samoyeds >= 0 && thisSue.samoyeds != giftDetail.samoyeds {
			continue
		}
		if thisSue.pomeranians >= 0 && thisSue.pomeranians >= giftDetail.pomeranians {
			continue
		}
		if thisSue.akitas >= 0 && thisSue.akitas != giftDetail.akitas {
			continue
		}
		if thisSue.vizslas >= 0 && thisSue.vizslas != giftDetail.vizslas {
			continue
		}
		if thisSue.goldfish >= 0 && thisSue.goldfish >= giftDetail.goldfish {
			continue
		}
		if thisSue.trees >= 0 && thisSue.trees <= giftDetail.trees {
			continue
		}
		if thisSue.cars >= 0 && thisSue.cars != giftDetail.cars {
			continue
		}
		if thisSue.perfumes >= 0 && thisSue.perfumes != giftDetail.perfumes {
			continue
		}

		return strconv.Itoa(sueIdx + 1)
	}
	return "no sue found"
}

type sue struct {
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func day16ExtractSues(in string) []sue {
	sues := []sue{}
	// Sue 1: children: 1, cars: 8, vizslas: 7
	sueDetails := utils.AllStringsFromRegex(in, `([a-z]+): ([0-9]+)`)
	for _, thisSueDetails := range sueDetails {

		thisSue := sue{
			children:    -1,
			cats:        -1,
			samoyeds:    -1,
			pomeranians: -1,
			akitas:      -1,
			vizslas:     -1,
			goldfish:    -1,
			trees:       -1,
			cars:        -1,
			perfumes:    -1,
		}
		for _, detail := range thisSueDetails {
			switch detail[1] {
			case "children":
				thisSue.children = utils.MustAtoi(detail[2])
			case "cats":
				thisSue.cats = utils.MustAtoi(detail[2])
			case "samoyeds":
				thisSue.samoyeds = utils.MustAtoi(detail[2])
			case "pomeranians":
				thisSue.pomeranians = utils.MustAtoi(detail[2])
			case "akitas":
				thisSue.akitas = utils.MustAtoi(detail[2])
			case "vizslas":
				thisSue.vizslas = utils.MustAtoi(detail[2])
			case "goldfish":
				thisSue.goldfish = utils.MustAtoi(detail[2])
			case "trees":
				thisSue.trees = utils.MustAtoi(detail[2])
			case "cars":
				thisSue.cars = utils.MustAtoi(detail[2])
			case "perfumes":
				thisSue.perfumes = utils.MustAtoi(detail[2])
			}
		}
		sues = append(sues, thisSue)
	}
	return sues
}

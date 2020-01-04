package days

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day13Part1(in string) string {
	// Alice would lose 2 happiness units by sitting next to David.
	// Bob would gain 83 happiness units by sitting next to Alice.
	happinessDetails := utils.StringsFromRegex(in, `^([a-zA-Z]+) would (lose|gain) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+)\.$`)
	happinessMap := map[string]map[string]int{}
	for _, details := range happinessDetails {
		attendee := details[1]

		curHappiness := happinessMap[attendee]
		if curHappiness == nil {
			curHappiness = map[string]int{}
		}
		if details[2] == "gain" {
			curHappiness[details[4]] = utils.MustAtoi(details[3])
		} else if details[2] == "lose" {
			curHappiness[details[4]] = -utils.MustAtoi(details[3])
		} else {
			panic(fmt.Errorf("unexpected happiness detail: %v", details))
		}
		happinessMap[attendee] = curHappiness
	}

	attendees := []string{}
	for attendee := range happinessMap {
		attendees = append(attendees, attendee)
	}
	attendeePermutations := utils.PermutationsString(attendees)
	maxHappiness := math.MinInt64
	for _, perm := range attendeePermutations {
		thisHappiness := 0

		for i := 1; i <= len(perm); i++ {
			prevAttendee := perm[(i-1)%len(perm)]
			attendee := perm[(i)%len(perm)]
			nextAttendee := perm[(i+1)%len(perm)]

			thisHappiness += happinessMap[attendee][prevAttendee]
			thisHappiness += happinessMap[attendee][nextAttendee]
		}

		maxHappiness = utils.Max(maxHappiness, thisHappiness)
	}

	return strconv.Itoa(maxHappiness)
}

func (r *Runner) Day13Part2(in string) string {
	// Alice would lose 2 happiness units by sitting next to David.
	// Bob would gain 83 happiness units by sitting next to Alice.
	happinessDetails := utils.StringsFromRegex(in, `^([a-zA-Z]+) would (lose|gain) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+)\.$`)
	happinessMap := map[string]map[string]int{}
	happinessMap["me"] = map[string]int{}
	for _, details := range happinessDetails {
		attendee := details[1]

		curHappiness := happinessMap[attendee]
		if curHappiness == nil {
			curHappiness = map[string]int{}
		}
		if details[2] == "gain" {
			curHappiness[details[4]] = utils.MustAtoi(details[3])
		} else if details[2] == "lose" {
			curHappiness[details[4]] = -utils.MustAtoi(details[3])
		} else {
			panic(fmt.Errorf("unexpected happiness detail: %v", details))
		}

		curHappiness["me"] = 0
		happinessMap[attendee] = curHappiness

		happinessMap["me"][attendee] = 0
	}

	attendees := []string{}
	for attendee := range happinessMap {
		attendees = append(attendees, attendee)
	}
	attendeePermutations := utils.PermutationsString(attendees)
	maxHappiness := math.MinInt64
	for _, perm := range attendeePermutations {
		thisHappiness := 0

		for i := 1; i <= len(perm); i++ {
			prevAttendee := perm[(i-1)%len(perm)]
			attendee := perm[(i)%len(perm)]
			nextAttendee := perm[(i+1)%len(perm)]

			thisHappiness += happinessMap[attendee][prevAttendee]
			thisHappiness += happinessMap[attendee][nextAttendee]
		}

		maxHappiness = utils.Max(maxHappiness, thisHappiness)
	}

	return strconv.Itoa(maxHappiness)
}

package days

import "github.com/fivegreenapples/AOC2015/utils"

import "math"

import "strconv"

type reindeerProperties struct {
	flightSpeed    int
	flightDuration int
	restDuration   int
}

const RaceTime = 2503

func (r *Runner) Day14Part1(in string) string {
	// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
	reindeerCapabilities := utils.StringsFromRegex(in, `^([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds\.$`)
	reindeer := map[string]reindeerProperties{}
	for _, rCap := range reindeerCapabilities {
		r := reindeerProperties{
			flightSpeed:    utils.MustAtoi(rCap[2]),
			flightDuration: utils.MustAtoi(rCap[3]),
			restDuration:   utils.MustAtoi(rCap[4]),
		}
		reindeer[rCap[1]] = r
	}

	maxDistance := math.MinInt64
	for _, properties := range reindeer {
		dist := day14DistanceAfterSeconds(properties, RaceTime)
		maxDistance = utils.Max(maxDistance, dist)
	}

	return strconv.Itoa(maxDistance)
}

func (r *Runner) Day14Part2(in string) string {
	// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
	reindeerCapabilities := utils.StringsFromRegex(in, `^([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds\.$`)
	reindeer := map[string]reindeerProperties{}
	for _, rCap := range reindeerCapabilities {
		r := reindeerProperties{
			flightSpeed:    utils.MustAtoi(rCap[2]),
			flightDuration: utils.MustAtoi(rCap[3]),
			restDuration:   utils.MustAtoi(rCap[4]),
		}
		reindeer[rCap[1]] = r
	}

	scores := map[string]int{}
	for s := 1; s <= RaceTime; s++ {

		maxDistance := math.MinInt64
		distances := map[string]int{}
		for name, properties := range reindeer {
			distance := day14DistanceAfterSeconds(properties, s)
			distances[name] = distance
			maxDistance = utils.Max(maxDistance, distance)
		}

		for name, distance := range distances {
			if distance == maxDistance {
				scores[name] += 1
			}
		}

	}

	maxScore := math.MinInt64
	for _, score := range scores {
		maxScore = utils.Max(maxScore, score)
	}

	return strconv.Itoa(maxScore)
}

func day14DistanceAfterSeconds(reindeer reindeerProperties, racetime int) int {
	dist := 0

	cycles := racetime / (reindeer.flightDuration + reindeer.restDuration)
	partCycleTime := racetime % (reindeer.flightDuration + reindeer.restDuration)

	// Add distance for complete flight+rest cycles achieved in race time
	dist += cycles * (reindeer.flightDuration * reindeer.flightSpeed)

	// Next add distance for the remaining time after the last full cycle
	partFlightTime := utils.Min(partCycleTime, reindeer.flightDuration)
	dist += (partFlightTime * reindeer.flightSpeed)

	return dist
}

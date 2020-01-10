package days

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func (r *Runner) Day22Part1(in string) string {
	me := d22PlayerStats{
		hitPoints: 50,
		mana:      500,
	}
	boss := d22PlayerStats{
		hitPoints: 51,
		damage:    9,
	}

	minCostOfWin := r.d22RunGame(me, boss, false)
	return strconv.Itoa(minCostOfWin)
}

func (r *Runner) Day22Part2(in string) string {
	me := d22PlayerStats{
		hitPoints: 50,
		mana:      500,
	}
	boss := d22PlayerStats{
		hitPoints: 51,
		damage:    9,
	}

	minCostOfWin := r.d22RunGame(me, boss, true)
	return strconv.Itoa(minCostOfWin)
}

func (r *Runner) d22RunGame(player, boss d22PlayerStats, hardMode bool) int {
	initialGameState := d22GameState{
		player:       player,
		boss:         boss,
		activeSpells: map[string]int{},
	}
	availableSpells := map[string]int{"recharge": 229, "shield": 113, "drain": 73, "poison": 173, "missile": 53}

	minCostOfWin := math.MaxInt64
	currentGameStates := []d22GameState{initialGameState}
	for len(currentGameStates) > 0 {

		if r.verbose {
			fmt.Println("Current num game states:", len(currentGameStates), "Current min cost:", minCostOfWin)
		}

		nextGameStates := []d22GameState{}
		for _, gs := range currentGameStates {

			for spell, cost := range availableSpells {

				newGS, won, err := gs.executeTurn(spell, cost, hardMode)
				if err != nil {
					// on error we died or invalid spell. move on
					continue
				}

				if newGS.totalSpend >= minCostOfWin {
					// if we've reached current lowest cost then we don't care
					// anymore for this permutation of spells
					continue
				}

				if won {
					// if we won then we know this is a lower win cost
					minCostOfWin = newGS.totalSpend
					continue
				}

				// otherwise add this newGS to stack of ones to keep trying with
				nextGameStates = append(nextGameStates, newGS)
			}

		}
		// Sort states by how weak the boss is so to get to a win state quicker
		sort.Slice(nextGameStates, func(i int, j int) bool {
			return nextGameStates[i].boss.hitPoints < nextGameStates[j].boss.hitPoints
		})

		currentGameStates = nextGameStates
	}

	return minCostOfWin
}

type d22PlayerStats struct {
	hitPoints int
	damage    int
	mana      int
}

type d22GameState struct {
	player       d22PlayerStats
	boss         d22PlayerStats
	activeSpells map[string]int
	totalSpend   int
}

func (current d22GameState) executeTurn(withSpell string, spellCost int, hardMode bool) (newState d22GameState, playerWon bool, problem error) {

	defer func() {
		if r := recover(); r != nil {
			problem = fmt.Errorf("%v", r)
		}
	}()

	newState = d22GameState{
		player:       current.player,
		boss:         current.boss,
		totalSpend:   current.totalSpend,
		activeSpells: map[string]int{},
	}

	if hardMode {
		newState.player.hitPoints--
		if newState.player.hitPoints <= 0 {
			return newState, false, errors.New("dead :-(")
		}
	}

	// player turn
	activeSpells := map[string]int{}
	for spell, timer := range current.activeSpells {
		switch spell {
		case "missile":
			panic("magic missile shouldn't be an active spell")
		case "drain":
			panic("drain shouldn't be an active spell")
		case "shield":
			// in effect but not relevant for player turn
		case "poison":
			newState.boss.hitPoints -= 3
		case "recharge":
			newState.player.mana += 101
		default:
			panic("not understood" + spell)
		}

		timer--
		if timer > 0 {
			activeSpells[spell] = timer
		}
	}

	// check if we won already
	if newState.boss.hitPoints <= 0 {
		return newState, true, nil
	}

	if _, active := activeSpells[withSpell]; active {
		return newState, false, errors.New("Can't use active spell")
	}

	newState.player.mana -= spellCost
	newState.totalSpend += spellCost
	// check if we're broke
	if newState.player.mana < 0 {
		return newState, false, errors.New("ran out of mana")
	}

	switch withSpell {
	case "missile":
		newState.boss.hitPoints -= 4
	case "drain":
		newState.boss.hitPoints -= 2
		newState.player.hitPoints += 2
	case "shield":
		activeSpells["shield"] = 6
	case "poison":
		activeSpells["poison"] = 6
	case "recharge":
		activeSpells["recharge"] = 5
	default:
		panic("not understood" + withSpell)
	}

	// boss turn
	playerArmor := 0
	for spell, timer := range activeSpells {
		switch spell {
		case "missile":
			panic("magic missile shouldn't be an active spell")
		case "drain":
			panic("drain shouldn't be an active spell")
		case "shield":
			playerArmor = 7
		case "poison":
			newState.boss.hitPoints -= 3
		case "recharge":
			newState.player.mana += 101
		default:
			panic("not understood" + spell)
		}

		timer--
		if timer > 0 {
			newState.activeSpells[spell] = timer
		}
	}

	// check if we won already
	if newState.boss.hitPoints <= 0 {
		return newState, true, nil
	}

	newState.player.hitPoints -= (newState.boss.damage - playerArmor)

	// check if we died
	if newState.player.hitPoints <= 0 {
		return newState, false, errors.New("dead :-(")
	}

	return newState, false, nil

}

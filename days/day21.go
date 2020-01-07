package days

import (
	"math"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day21Part1(in string) string {

	// 1 weapon, optional armor, up to two rings, rings can't be the same

	playerCombos := []d21ShoppingList{}
	for _, weapon := range d21Weapons {
		for _, armor := range d21Armor {
			for firstRing, ring1 := range d21Rings {
				for secondRing, ring2 := range d21Rings {

					if secondRing != "Plastic" && secondRing == firstRing {
						continue
					}

					list := d21ShoppingList{
						weapon,
						armor,
						ring1,
						ring2,
					}
					playerCombos = append(playerCombos, list)
				}
			}
		}
	}

	minCost := math.MaxInt64
	for _, c := range playerCombos {
		player := d21Character{
			hitPoints: 100,
			damage:    c.damage(),
			armor:     c.armor(),
		}
		if player.isWinnerAgainst(d21Boss) && c.cost() < minCost {
			minCost = c.cost()

		}
	}

	return strconv.Itoa(minCost)
}

func (r *Runner) Day21Part2(in string) string {
	// 1 weapon, optional armor, up to two rings, rings can't be the same

	playerCombos := []d21ShoppingList{}
	for _, weapon := range d21Weapons {
		for _, armor := range d21Armor {
			for firstRing, ring1 := range d21Rings {
				for secondRing, ring2 := range d21Rings {

					if secondRing != "Plastic" && secondRing == firstRing {
						continue
					}

					list := d21ShoppingList{
						weapon,
						armor,
						ring1,
						ring2,
					}
					playerCombos = append(playerCombos, list)
				}
			}
		}
	}

	maxCost := math.MinInt64
	for _, c := range playerCombos {
		player := d21Character{
			hitPoints: 100,
			damage:    c.damage(),
			armor:     c.armor(),
		}
		if !player.isWinnerAgainst(d21Boss) && c.cost() > maxCost {
			maxCost = c.cost()

		}
	}

	return strconv.Itoa(maxCost)
}

type d21Character struct {
	hitPoints int
	damage    int
	armor     int
}

type d21Item struct {
	cost   int
	damage int
	armor  int
}

type d21ShoppingList []d21Item

func (player d21Character) isWinnerAgainst(enemy d21Character) bool {
	playerOnEnemyDamage := utils.Max(player.damage-enemy.armor, 1)
	enemyOnPlayerDamage := utils.Max(enemy.damage-player.armor, 1)

	enemyRoundsToDie := 1 + enemy.hitPoints/playerOnEnemyDamage
	playerRoundsToDie := 1 + player.hitPoints/enemyOnPlayerDamage

	// player goes first so equal rounds means player wins
	return playerRoundsToDie >= enemyRoundsToDie
}

func (list d21ShoppingList) cost() int {
	val := 0
	for _, i := range list {
		val += i.cost
	}
	return val
}
func (list d21ShoppingList) armor() int {
	val := 0
	for _, i := range list {
		val += i.armor
	}
	return val
}
func (list d21ShoppingList) damage() int {
	val := 0
	for _, i := range list {
		val += i.damage
	}
	return val
}

// these details are from puzzle input
var d21Boss = d21Character{
	hitPoints: 109,
	damage:    8,
	armor:     2,
}

var d21Weapons = map[string]d21Item{
	"Dagger":     {cost: 8, damage: 4},
	"Shortsword": {cost: 10, damage: 5},
	"Warhammer":  {cost: 25, damage: 6},
	"Longsword":  {cost: 40, damage: 7},
	"Greataxe":   {cost: 74, damage: 8},
}

var d21Armor = map[string]d21Item{
	"BirthdaySuit": {cost: 0},
	"Leather":      {cost: 13, armor: 1},
	"Chainmail":    {cost: 31, armor: 2},
	"Splintmail":   {cost: 53, armor: 3},
	"Bandedmail":   {cost: 75, armor: 4},
	"Platemail":    {cost: 102, armor: 5},
}

var d21Rings = map[string]d21Item{
	"Plastic":    {cost: 0},
	"Damage +1":  {cost: 25, damage: 1},
	"Damage +2":  {cost: 50, damage: 2},
	"Damage +3":  {cost: 100, damage: 3},
	"Defense +1": {cost: 20, armor: 1},
	"Defense +2": {cost: 40, armor: 2},
	"Defense +3": {cost: 80, armor: 3},
}

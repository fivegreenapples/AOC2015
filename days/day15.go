package days

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day15Part1(in string) string {
	ingredients := day15ProcessIngredientDetails(in)

	maxRecipeScore := math.MinInt64
	allRecipeOptions := day15FindRecipeOptions(ingredients, 100)
	for _, recipeOption := range allRecipeOptions {
		maxRecipeScore = utils.Max(maxRecipeScore, recipeOption.score())
		if r.verbose {
			fmt.Println(recipeOption, recipeOption.score())
		}
	}

	return strconv.Itoa(maxRecipeScore)
}

func (r *Runner) Day15Part2(in string) string {
	ingredients := day15ProcessIngredientDetails(in)

	maxRecipeScore := math.MinInt64
	allRecipeOptions := day15FindRecipeOptions(ingredients, 100)
	for _, recipeOption := range allRecipeOptions {
		if recipeOption.calories() == 500 {
			maxRecipeScore = utils.Max(maxRecipeScore, recipeOption.score())
			if r.verbose {
				fmt.Println(recipeOption, recipeOption.score())
			}
		}
	}

	return strconv.Itoa(maxRecipeScore)
}

type ingredientProperties struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func (i ingredientProperties) String() string {
	// return fmt.Sprintf("%s - c:%d d:%d f:%d t:%d", i.name, i.capacity, i.durability, i.flavor, i.texture)
	return i.name
}
func (i ingredientProperties) score() int {
	if i.capacity <= 0 || i.durability <= 0 || i.flavor <= 0 || i.texture <= 0 {
		return 0
	}
	return i.capacity * i.durability * i.flavor * i.texture
}

type recipeItem struct {
	ingredient ingredientProperties
	teaspoons  int
}

func (r recipeItem) String() string {
	return fmt.Sprintf("%s:%d", r.ingredient, r.teaspoons)
}

type recipe []recipeItem

func (r recipe) score() int {
	score := ingredientProperties{}
	for _, item := range r {
		score.capacity += (item.ingredient.capacity * item.teaspoons)
		score.durability += (item.ingredient.durability * item.teaspoons)
		score.flavor += (item.ingredient.flavor * item.teaspoons)
		score.texture += (item.ingredient.texture * item.teaspoons)
	}
	return score.score()
}

func (r recipe) calories() int {
	calories := 0
	for _, item := range r {
		calories += (item.ingredient.calories * item.teaspoons)
	}
	return calories
}

func day15ProcessIngredientDetails(in string) []ingredientProperties {
	// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	// Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
	ingredientDetails := utils.StringsFromRegex(in, `^([a-zA-Z]+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)$`)
	ingredients := []ingredientProperties{}
	for _, details := range ingredientDetails {
		ingredients = append(ingredients, ingredientProperties{
			name:       details[1],
			capacity:   utils.MustAtoi(details[2]),
			durability: utils.MustAtoi(details[3]),
			flavor:     utils.MustAtoi(details[4]),
			texture:    utils.MustAtoi(details[5]),
			calories:   utils.MustAtoi(details[6]),
		})
	}
	return ingredients
}

func day15FindRecipeOptions(ingredients []ingredientProperties, remainingTeaspoons int) []recipe {
	if len(ingredients) == 1 {
		return []recipe{
			[]recipeItem{
				{
					ingredient: ingredients[0],
					teaspoons:  remainingTeaspoons,
				},
			},
		}
	}

	recipes := []recipe{}

	thisIngredient := ingredients[0]
	for t := 0; t <= remainingTeaspoons; t++ {
		thisRecipeItem := recipeItem{
			ingredient: thisIngredient,
			teaspoons:  t,
		}
		recipeOptions := day15FindRecipeOptions(ingredients[1:], remainingTeaspoons-t)

		for _, option := range recipeOptions {

			thisRecipe := recipe{
				thisRecipeItem,
			}
			thisRecipe = append(thisRecipe, option...)

			recipes = append(recipes, thisRecipe)
		}

	}
	return recipes
}

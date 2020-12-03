package _2018

import (
	"bytes"
	"fmt"
)

const RecipesToIdentify = 10

type elf struct {
	recipe, position int
}

func (e *elf) findNextRecipe(recipes []int) {
	e.position = (e.position + e.recipe + 1) % len(recipes)
	e.recipe = recipes[e.position]
}

type lab struct {
	first, second *elf
	recipes       []int
}

func newLab(first, second *elf) *lab {
	first.position = 0
	second.position = 1

	return &lab{first, second, []int{first.recipe, second.recipe}}
}

func (l *lab) work() {
	newRecipe := l.first.recipe + l.second.recipe
	if newRecipe < 10 {
		l.recipes = append(l.recipes, newRecipe)
	} else {
		newRecipes := []int{newRecipe / 10 % 10, newRecipe % 10}
		l.recipes = append(l.recipes, newRecipes...)
	}

	l.first.findNextRecipe(l.recipes)
	l.second.findNextRecipe(l.recipes)
}

func (l *lab) bestTenRecipesAfter(n int) string {
	for len(l.recipes) < n+RecipesToIdentify {
		l.work()
	}

	var b bytes.Buffer
	for i := 0; i < RecipesToIdentify; i++ {
		fmt.Fprintf(&b, "%d", l.recipes[i+n])
	}

	return b.String()
}

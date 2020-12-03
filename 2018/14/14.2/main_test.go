package _2018

import (
	"reflect"
	"testing"
)

func TestRecipesAreCreatedAddingTwoOfThem(t *testing.T) {
	first := &elf{1, -1}
	second := &elf{3, -1}

	l := newLab(first, second)
	l.work()

	if r := l.recipes; !reflect.DeepEqual(r, []int{1, 3, 4}) {
		t.Errorf("recipes should be {1,3,4}, got %+v\n", r)
	}
}

func TestCreatingRecipesThanHaveACombinationGreaterThanTen(t *testing.T) {
	first := &elf{7, -1}
	second := &elf{5, -1}

	l := newLab(first, second)
	l.work()

	if r := l.recipes; !reflect.DeepEqual(r, []int{7, 5, 1, 2}) {
		t.Errorf("recipes should be {7,5,1,2}, got %+v\n", r)
	}
}

func TestElfsMoveToTheNextRecipeAfterTheyAreDoneWithTheirCurrentOne(t *testing.T) {
	first := &elf{1, -1}
	second := &elf{2, -1}

	l := newLab(first, second)
	l.work()

	if first.position != 1 {
		t.Errorf("first elf should be in second position because she just worked on a recipe with value 1, got %d\n", first.position)
	}

	if second.position != 0 {
		t.Errorf("first elf should be in position first because she just worked on a recipe with value 2, got %d\n", second.position)
	}
}

func TestCreatingRecipes(t *testing.T) {
	type testCase struct {
		nRecipes int
		expectedRecipe string
	}

	tcs := []testCase{
		{
			5, "0124515891",
		},{
			9, "5158916779",
		},{
			18, "9251071085",
		},{
			2018, "5941429882",
		},{
			360781, "6521571010",
		},
	}

	for _, tc := range tcs {
		first := &elf{3, -1}
		second := &elf{7, -1}

		l := newLab(first, second)
		if b := l.bestTenRecipesAfter(tc.nRecipes); b != tc.expectedRecipe {
			t.Errorf("failed getting best after %d, got %s but expecting %s", tc.nRecipes, b, tc.expectedRecipe)
		}

	}
}
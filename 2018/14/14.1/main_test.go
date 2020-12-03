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

func TestCountingRecipes(t *testing.T) {
	type testCase struct {
		searchedRecipe   string
		expectedNRecipes int
	}

	tcs := []testCase{
		{
			"01245", 5,
		}, {
			"51589", 9,
		}, {
			"92510", 18,
		}, {
			"59414", 2018,
		}, {
			"360781", 20262967,
		},
	}

	for _, tc := range tcs {
		first := &elf{3, -1}
		second := &elf{7, -1}

		l := newLab(first, second)
		if b := l.countRecipesUntil(tc.searchedRecipe); b != tc.expectedNRecipes {
			t.Errorf("failed getting counting recipes until %s, got %d but expecting %d", tc.searchedRecipe, b, tc.expectedNRecipes)
		}

	}
}

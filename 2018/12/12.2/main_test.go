package _2018

import (
	"testing"
)

func TestItCanDetectARuleThatAppliesToThePotInTheMiddle(t *testing.T) {
	rules := map[string]string{}
	rules[".#..."] = "#"

	g := newGarden(".#...", rules)
	g.growGeneration()

	if p := g.potsAsString(); p != ".##.." {
		t.Errorf("Expecting pots to become .##.. after one generation, got %s\n", p)
	}
}

func TestItCanDetectARuleThatAppliesToThePotBeforeTheFirstOne(t *testing.T) {
	rules := map[string]string{}
	rules["...#."] = "#"

	g := newGarden("#....", rules)
	g.growGeneration()

	expected := "##...."

	if p := g.potsAsString(); p != expected {
		t.Errorf("Expecting pots to become %s after one generation, got %s\n", expected, p)
	}
}

func TestItCanDetectARuleThatAppliesToThePotAfterTheLastOne(t *testing.T) {
	rules := map[string]string{}
	rules["#...."] = "#"

	g := newGarden("...#.", rules)
	g.growGeneration()
	expected := "...#.#"

	if p := g.potsAsString(); p != expected {
		t.Errorf("Expecting pots to become %s after one generation, got %s\n", expected, p)
	}
}

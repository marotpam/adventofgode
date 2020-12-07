package _2020

import (
	"strconv"
	"strings"
)

type bagCollection map[string]bag

func newBagCollection(bags []bag) bagCollection {
	collection := make(map[string]bag, len(bags))
	for _, b := range bags {
		collection[b.description] = b
	}
	return collection
}

type innerBag struct {
	description string
	quantity    int
}

type bag struct {
	description string
	innerBags   []innerBag
}

func CountBagsContainingShinyGoldBag(rawInput string) int {
	collection := newBagCollection(parse(strings.Split(rawInput, ".\n")))

	c := 0
	for _, b := range parse(strings.Split(rawInput, ".\n")) {
		if collection.containsShinyGoldBag(b) {
			c++
		}
	}

	return c
}

func (c bagCollection) containsShinyGoldBag(b bag) bool {
	for _, i := range b.innerBags {
		if i.description == "shiny gold" || c.containsShinyGoldBag(c[i.description]) {
			return true
		}
	}

	return false
}

func parse(lines []string) []bag {
	bags := make([]bag, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, " bags contain ")

		rawInnerBags := strings.Split(parts[1], ", ")
		innerBags := make([]innerBag, 0, len(rawInnerBags))
		for _, rawInnerBag := range rawInnerBags {
			if rawInnerBag == "no other bags" {
				continue
			}
			parts := strings.Split(rawInnerBag, " ")
			q, _ := strconv.Atoi(parts[0])
			innerBags = append(innerBags, innerBag{
				description: parts[1] + " " + parts[2],
				quantity:    q,
			})
		}

		bags = append(bags, bag{
			description: parts[0],
			innerBags:   innerBags,
		})
	}
	return bags
}

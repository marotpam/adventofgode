package _2020

import "strings"

func CountIndividualBags(rawInput string) int {
	collection := newBagCollection(parse(strings.Split(rawInput, ".\n")))

	return collection.countInnerBags(collection["shiny gold"]) - 1
}

func (c bagCollection) countInnerBags(b bag) int {
	count := 1
	for _, i := range b.innerBags {
		count += i.quantity * c.countInnerBags(c[i.description])
	}
	return count
}

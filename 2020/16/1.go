package _2020

import (
	"strconv"
	"strings"
)

func CalculateTicketScanningErrorRate(rawInput string) int {
	rules, _, nearbyTickets := parse(rawInput)

	c := 0
	for _, t := range nearbyTickets {
		c += t.errorRate(rules)
	}
	return c
}

type intRange struct {
	lo, hi int
}

func (r intRange) includes(n int) bool {
	return r.lo <= n && n <= r.hi
}

type rule struct {
	name   string
	ranges []intRange
}

func (r rule) includes(n int) bool {
	for _, intRange := range r.ranges {
		if intRange.includes(n) {
			return true
		}
	}
	return false
}

type ticket struct {
	fields []int
}

func (t ticket) errorRate(rules []rule) int {
	errorRate := 0
	for _, f := range t.fields {
		hasRule := false
		for _, r := range rules {
			if r.includes(f) {
				hasRule = true
				break
			}
		}
		if !hasRule {
			errorRate += f
		}
	}
	return errorRate
}

func parse(input string) ([]rule, ticket, []ticket) {
	parts := strings.Split(input, "\n\n")

	rules := make([]rule, 0, len(parts[0]))
	for _, l := range strings.Split(parts[0], "\n") {
		ruleParts := strings.Split(l, ": ")
		rawRanges := strings.Split(ruleParts[1], " or ")
		ranges := make([]intRange, 0, len(rawRanges))
		for _, r := range rawRanges {
			ranges = append(ranges, parseIntRange(r))
		}
		rules = append(rules, rule{
			name:   ruleParts[0],
			ranges: ranges,
		})
	}

	myTicketParts := strings.Split(parts[1], "\n")
	myTicket := parseTicket(myTicketParts[1])

	nearbyTicketParts := strings.Split(parts[2], "\n")
	nearbyTickets := make([]ticket, 0, len(nearbyTicketParts) - 1)
	for _, part := range nearbyTicketParts[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(part))
	}

	return rules, myTicket, nearbyTickets
}

func parseIntRange(rawRange string) intRange {
	rangeParts := strings.Split(rawRange, "-")
	return intRange{
		lo: num(rangeParts[0]),
		hi: num(rangeParts[1]),
	}
}

func num(raw string) int {
	n, _ := strconv.Atoi(raw)
	return n
}

func parseTicket(rawTicket string) ticket {
	parts := strings.Split(rawTicket, ",")
	fields := make([]int, 0, len(parts))

	for _, part := range parts {
		fields = append(fields, num(part))
	}

	return ticket{fields: fields}
}
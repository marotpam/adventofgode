package _2020

import (
	"strconv"
	"strings"
)

var emptyMap = map[string]bool{"": true}

func CountMessagesMatchingRules(rawInput string) int {
	rules, messages := parse(rawInput)
	all := rules.generateAll()

	c := 0
	for _, m := range messages {
		_, ok := all[m]
		if ok {
			c++
		}
	}
	return c
}

type rule struct {
	exactMatch string
	innerRules [][]int
}

type rules map[int]rule

func (rs rules) generateAll() map[string]bool {
	return rs.generate(rs[0], emptyMap)
}

func (rs rules) generate(rule rule, previous map[string]bool) map[string]bool {
	if rule.exactMatch != "" {
		return appendToAllKeys(rule.exactMatch, previous)
	}

	res := map[string]bool{}
	for _, innerRule := range rule.innerRules {
		m := emptyMap
		for _, r := range innerRule {
			generate := rs.generate(rs[r], previous)
			m = appendMapKeysToAllKeys(m, generate)
		}
		res = merge(res, m)
	}

	return res
}

func appendMapKeysToAllKeys(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool, len(a))
	for k1 := range a {
		for k2 := range b {
			res[k1+k2] = true
		}
	}
	return res
}

func merge(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool, len(a)+len(b))
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}

func appendToAllKeys(suffix string, m map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k, v := range m {
		res[k+suffix] = v
	}
	return res
}

func parse(input string) (rules, []string) {
	parts := strings.Split(input, "\n\n")

	rawRules := strings.Split(parts[0], "\n")
	rules := make(map[int]rule, len(rawRules))
	for _, l := range rawRules {
		lineParts := strings.Split(l, ": ")

		if strings.Contains(l, "\"") {
			rules[parseNumber(lineParts[0])] = rule{
				exactMatch: strings.Trim(lineParts[1], "\""),
			}
			continue
		}

		rawInnerRules := strings.Split(lineParts[1], " | ")
		innerRules := make([][]int, 0, len(rawInnerRules))
		for _, innerRule := range rawInnerRules {
			innerRules = append(innerRules, parseNumbers(innerRule))
		}

		rules[parseNumber(lineParts[0])] = rule{
			innerRules: innerRules,
		}
	}

	return rules, strings.Split(parts[1], "\n")
}

func parseNumbers(innerRule string) []int {
	ints := []int{}
	for _, i := range strings.Split(innerRule, " ") {
		ints = append(ints, parseNumber(i))
	}
	return ints
}

func parseNumber(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

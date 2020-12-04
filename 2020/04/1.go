package _2020

import (
	"strings"
)

type passport map[string]string

func (p passport) isValid1() bool {
	required := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, r := range required {
		if p[r] == "" {
			return false
		}
	}

	return true
}

func CountValidPasswords1(rawInput string) int {
	c := 0

	passports := strings.Split(rawInput, "\n\n")
	for _, p := range passports {
		rawFields := strings.FieldsFunc(p, func(r rune) bool {
			return r == '\n' || r == ' '
		})

		pass := make(passport)
		for _, f := range rawFields {
			parts := strings.Split(f, ":")
			_, ok := pass[parts[0]]
			if ok {
				break
			}
			pass[parts[0]] = parts[1]
		}

		if pass.isValid1() {
			c++
		}
	}

	return c
}

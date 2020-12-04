package _2020

import (
	"regexp"
	"strconv"
	"strings"
)

type validationFunc func(string) bool

func (p passport) isValid2() bool {
	required := map[string]validationFunc{
		"byr": p.validateBirthYear,
		"iyr": p.validateIssueYear,
		"eyr": p.validateExpirationYear,
		"hgt": p.validateHeight,
		"hcl": p.validateHairColour,
		"ecl": p.validateEyeColour,
		"pid": p.validatePassportID,
	}
	for k, validationFunc := range required {
		if p[k] == "" || !validationFunc(p[k]) {
			return false
		}
	}

	return true
}

func (p passport) validateBirthYear(s string) bool {
	return isNumberInRange(s, 1920, 2002)
}

func (p passport) validateIssueYear(s string) bool {
	return isNumberInRange(s, 2010, 2020)
}

func (p passport) validateExpirationYear(s string) bool {
	return isNumberInRange(s, 2020, 2030)
}

func isNumberInRange(s string, lo int, hi int) bool {
	n, err := strconv.Atoi(s)

	return err == nil && n >= lo && n <= hi
}

func (p passport) validateHeight(s string) bool {
	n, system := s[0:len(s)-2], s[len(s)-2:]

	return (system == "in" && isNumberInRange(n, 59, 76)) ||
		(system == "cm" && isNumberInRange(n, 150, 193))
}

func (p passport) validateHairColour(s string) bool {
	return matchesPattern(s, "^#[0-9a-f]{6}$")
}

func (p passport) validateEyeColour(s string) bool {
	return matchesPattern(s, "^amb|blu|brn|gry|grn|hzl|oth$")
}

func (p passport) validatePassportID(s string) bool {
	return matchesPattern(s, "^[0-9]{9}$")
}

func matchesPattern(s string, pattern string) bool {
	matchString, err := regexp.MatchString(pattern, s)
	return err == nil && matchString
}

func CountValidPasswords2(rawInput string) int {
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

		if pass.isValid2() {
			c++
		}
	}

	return c
}

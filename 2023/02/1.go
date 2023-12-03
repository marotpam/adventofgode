package _2023

import (
	"strconv"
	"strings"
)

type game []set

func (g game) canBePlayed() bool {
	for _, pick := range g {
		scores := map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}
		for _, cubePick := range pick {
			if scores[cubePick.colour] < cubePick.count {
				return false
			}
		}
	}
	return true
}

type set []cubePick

type cubePick struct {
	colour string
	count  int
}

func SumIdsOfPossibleGames(rawInput string) int {
	sum := 0
	for id, game := range parseGames(rawInput) {
		if game.canBePlayed() {
			sum += (id + 1)
		}
	}

	return sum
}

func parseGames(rawInput string) []game {
	rawGames := strings.Split(rawInput, "\n")
	games := make([]game, 0, len(rawGames))

	for _, rg := range rawGames {
		gameParts := strings.Split(rg, ": ")
		if len(gameParts) != 2 {
			continue
		}
		rawSets := strings.Split(gameParts[1], "; ")
		sets := make([]set, 0, len(rawSets))
		for _, rs := range rawSets {
			rawCubes := strings.Split(rs, ", ")
			cubes := make([]cubePick, 0, len(rawCubes))
			for _, rc := range rawCubes {
				cubeParts := strings.Split(rc, " ")
				count, _ := strconv.Atoi(cubeParts[0])
				cubes = append(cubes, cubePick{
					count:  count,
					colour: cubeParts[1],
				})
			}
			sets = append(sets, cubes)
		}
		games = append(games, sets)
	}
	return games
}

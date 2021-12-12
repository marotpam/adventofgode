package _2021

import (
	"strconv"
	"strings"
)

func SumRiskLevels(rawInput string) int {
	hm := parseHeatMap(rawInput)
	sum := 0
	for _, h := range hm.getLowPointHeights() {
		sum += h + 1
	}
	return sum
}

type position struct {
	x, y int
}

type heatMap struct {
	locations     map[position]int
	width, height int
}

func (m heatMap) getLowPointHeights() []int {
	heights := make([]int, 0, 0)
	for pos, height := range m.locations {
		if m.isLowPoint(pos, height) {
			heights = append(heights, height)
		}
	}
	return heights
}

func (m heatMap) isLowPoint(pos position, height int) bool {
	for _, h := range m.findNeighbourHeights(pos) {
		if h <= height {
			return false
		}
	}
	return true
}

func (m heatMap) findNeighbourHeights(pos position) []int {
	tentative := []position{
		{x: pos.x, y: pos.y + 1},
		{x: pos.x, y: pos.y - 1},
		{x: pos.x + 1, y: pos.y},
		{x: pos.x - 1, y: pos.y},
	}
	heights := make([]int, 0, len(tentative))
	for _, p := range tentative {
		height, ok := m.locations[p]
		if ok {
			heights = append(heights, height)
		}
	}
	return heights
}

func parseHeatMap(rawInput string) heatMap {
	lines := strings.Split(rawInput, "\n")
	width := 0
	if len(lines) > 0 {
		width = len(lines[0])
	}

	locations := map[position]int{}
	for y := 0; y < len(lines); y++ {
		l := lines[y]
		for x := 0; x < len(l); x++ {
			n, _ := strconv.Atoi(string(l[x]))
			locations[position{
				x: y,
				y: x,
			}] = n
		}
	}
	return heatMap{locations: locations, width: width, height: len(lines)}
}

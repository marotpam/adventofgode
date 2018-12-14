package main

type coordinate struct {
	x, y int
}

type grid struct {
	points [][]int
}

func newGrid(ps []coordinate) *grid {
	points := [][]int{}
	row := []int{}

	for y := 0; y < 4; y++ {
		row = append(row, -1)
	}
	points = append(points, row)

	for i, p := range ps {
		points[p.x][p.y] = i
	}

	return &grid{points}
}

func (g *grid) closestPoints() [][]int {
	r := make([][]int, len(g.points))

	for i := 0; i < len(g.points); i++ {
		for j := 0; j < len(g.points[i]); j++ {
			r[i][j] = g.closestPoint(i, j)
		}
	}

	return r
}

func (g *grid) closestPoint(x, y int) int {
	for l := 0; x - l >= 0 && x + l <= len(g.points) && y - l > 0 && y + l <= len(g.points); l++ {

	}
	return 0
}
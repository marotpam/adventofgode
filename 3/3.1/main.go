package main

import (
	"bufio"
	"fmt"
	"os"
)

type claim struct {
	left, top, wide, tall int
}

func readClaims() []claim {
	c := []claim{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		var number, left, top, wide, tall int
		fmt.Sscanf(fileScanner.Text(), `#%d @ %d,%d: %dx%d`, &number, &left, &top, &wide, &tall)

		c = append(c, claim{left, top, wide, tall})
	}

	return c
}

type fabric struct {
	layout [][]int
}

func newFabric(cs []claim) fabric {
	x, y := getDimensions(cs)

	layout := matrixOfSize(x, y)

	return fabric{applyClaims(layout, cs)}
}

func (f *fabric) countOverlapping() int {
	count := 0
	for i := 0; i < len(f.layout); i++ {
		for j := 0; j < len(f.layout[0]); j++ {
			if f.layout[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func applyClaims(m [][]int, cs []claim) [][]int {
	withClaims := m

	for _, c := range cs {
		for i := 0; i < c.tall; i++ {
			for j := 0; j < c.wide; j++ {
				withClaims[c.top+i][c.left+j]++
			}
		}
	}

	return withClaims
}

func matrixOfSize(sizeX, sizeY int) [][]int {
	matrix := make([][]int, sizeY)

	for i:= 0; i < sizeY; i++ {
		matrix[i] = make([]int, sizeX)
	}

	return matrix
}

func getDimensions(cs []claim) (int, int) {
	maxX, maxY := 0, 0
	for _, c := range cs {
		x := c.left + c.wide
		if x > maxX {
			maxX = x
		}

		y:= c.top + c.tall
		if y > maxY {
			maxY = y
		}
	}

	return maxX, maxY
}

func main() {
	f := newFabric(readClaims())

	fmt.Println(f.countOverlapping()) // 116920
}

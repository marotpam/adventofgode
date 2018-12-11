package main

import (
	"bufio"
	"fmt"
	"os"
)

type claim struct {
	id, left, top, wide, tall int
}

func readClaims() []claim {
	c := []claim{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		var id, left, top, wide, tall int
		fmt.Sscanf(fileScanner.Text(), `#%d @ %d,%d: %dx%d`, &id, &left, &top, &wide, &tall)

		c = append(c, claim{id, left, top, wide, tall})
	}

	return c
}

type position struct {
	cs []int
}

type fabric struct {
	layout [][]position
	overlappingClaims map[int]bool
}

func newFabric(cs []claim) fabric {
	x, y := getDimensions(cs)

	layout := matrixOfSize(x, y)

	l, o := applyClaims(layout, cs)

	return fabric{l, 	o}
}

func (f *fabric) nonOverlappingClaim() int {
	res := -1
	for id, overlaps := range f.overlappingClaims {
		fmt.Println(id, overlaps)
		if !overlaps {
			res = id
			fmt.Printf("%d doesnt overlap\n", res)
		}
	}

	return res
}

func (f *fabric) countOverlapping() int {
	count := 0
	for i := 0; i < len(f.layout); i++ {
		for j := 0; j < len(f.layout[0]); j++ {
			if len(f.layout[i][j].cs) > 1 {
				count++
			}
		}
	}
	return count
}

func applyClaims(m [][]position, cs []claim) ([][]position, map[int]bool) {
	withClaims := m
	overlapping := map[int]bool{}

	for _, c := range cs {
		for i := 0; i < c.tall; i++ {
			for j := 0; j < c.wide; j++ {
				withClaims[c.top+i][c.left+j].cs = append(withClaims[c.top+i][c.left+j].cs, c.id)
				if len(withClaims[c.top+i][c.left+j].cs) > 1 {
					for _, o := range withClaims[c.top+i][c.left+j].cs {
						overlapping[o] = true
					}
				} else {
					overlapping[c.id] = false
				}
			}
		}
	}

	return withClaims, overlapping
}

func matrixOfSize(sizeX, sizeY int) [][]position {
	matrix := make([][]position, sizeY)

	for i:= 0; i < sizeY; i++ {
		matrix[i] = make([]position, sizeX)
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

	//fmt.Println(f.nonOverlappingClaim())
	fmt.Println(f.countOverlapping())
}

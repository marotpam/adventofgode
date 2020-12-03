package _2018

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
)

type xy struct{ x, y int }

type velocity struct{ vx, vy int }

type point struct {
	position xy
	velocity velocity
}

func (p *point) isInPosition(pos xy) bool {
	return reflect.DeepEqual(p.position, pos)
}

type sky struct {
	maxX, maxY int
	minX, minY int
	points     []*point
}

func (s *sky) hasClearMessage() bool {
	return (s.maxY - s.minY) < 10
}

func (s *sky) hasPointInPos(xy xy) bool {
	for _, p := range s.points {
		if p.isInPosition(xy) {
			return true
		}
	}
	return false
}

func (s *sky) print() string {
	res := []string{}
	for y := s.minY; y <= s.maxY; y++ {
		row := []string{}
		for x := s.minX; x <= s.maxX; x++ {
			c := "."
			if s.hasPointInPos(xy{x, y}) {
				c = "#"
			}
			if x == 0 && y == 0 {
				c = "0"
			}
			row = append(row, c)
		}
		res = append(res, strings.Join(row, ""))
	}
	return strings.Join(res, "\n")
}

func (s *sky) tick() {
	maxX, maxY := math.MinInt64, math.MinInt64
	minX, minY := math.MaxInt64, math.MaxInt64

	for _, p := range s.points {
		p.position.x += p.velocity.vx
		p.position.y += p.velocity.vy

		if p.position.x > maxX {
			maxX = p.position.x
		}
		if p.position.y > maxY {
			maxY = p.position.y
		}
		if p.position.x < minX {
			minX = p.position.x
		}
		if p.position.y < minY {
			minY = p.position.y
		}
	}

	s.maxX = maxX
	s.maxY = maxY
	s.minX = minX
	s.minY = minY
}

func readPoints() []*point {
	c := []*point{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		var x, y, vx, vy int
		fmt.Sscanf(fileScanner.Text(), `position=<%d, %d> velocity=<%d, %d>`, &x, &y, &vx, &vy)

		xy := xy{x, y}
		v := velocity{vx, vy}
		p := &point{
			position: xy,
			velocity: v,
		}

		c = append(c, p)
	}

	return c
}

func main() {
	s := sky{math.MaxInt64,math.MaxInt64, 0, 0, readPoints()}

	var i int
	for i = 0; !s.hasClearMessage(); i++ {
		s.tick()
	}

	fmt.Println(i) // 10240
	fmt.Print(s.print()) // RLEZNRAN
}

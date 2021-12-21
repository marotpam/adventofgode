package _2021

import (
	"strconv"
	"strings"
)

func CountDotsAfterFoldCount(rawInput string, foldCount int) int {
	dots, foldInstructions := parse(rawInput)
	paper := newPaper(dots)
	for i := 0; i < foldCount; i++ {
		paper.fold(foldInstructions[i])
	}
	return paper.countVisible()
}

type paper struct {
	dots   map[position]bool
	width  int
	height int
}

func newPaper(dots []position) paper {
	width, height := 0, 0
	dotMap := make(map[position]bool, len(dots))
	for _, dot := range dots {
		dotMap[dot] = true
		if dot.y > height {
			height = dot.y
		}
		if dot.x > width {
			width = dot.x
		}
	}
	return paper{
		height: height,
		width:  width,
		dots:   dotMap,
	}
}

func (p *paper) countVisible() int {
	return len(p.dots)
}

func (p *paper) fold(inst instruction) {
	startX, startY := 0, 0
	newWidth, newHeight := p.width, p.height

	switch inst.axis {
	case 'x':
		startX = inst.index + 1
		newWidth = newWidth / 2
	case 'y':
		startY = inst.index + 1
		newHeight = newHeight / 2
	}

	for y := startY; y <= p.height; y++ {
		for x := startX; x <= p.width; x++ {
			pos := position{
				x: x,
				y: y,
			}

			mirrorPos := pos.mirror(inst)
			mirrorDot := p.dots[mirrorPos]
			if !mirrorDot {
				p.dots[mirrorPos] = p.dots[pos]
			}
			delete(p.dots, pos)
		}

	}

	p.width = newWidth
	p.height = newHeight
}

type position struct {
	x, y int
}

func (p position) mirror(inst instruction) position {
	mirror := position{
		x: p.x,
		y: p.y,
	}
	switch inst.axis {
	case 'x':
		mirror.x = inst.index - (p.x - inst.index)
	case 'y':
		mirror.y = inst.index - (p.y - inst.index)
	}
	return mirror
}

type instruction struct {
	axis  rune
	index int
}

func parse(input string) ([]position, []instruction) {
	var dots []position

	i := 0
	lines := strings.Split(input, "\n")
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			i++
			break
		}

		dots = append(dots, parsePosition(lines[i]))
	}

	var instructions []instruction
	for ; i < len(lines); i++ {
		instructions = append(instructions, parseInstruction(lines[i]))
	}

	return dots, instructions
}

func parsePosition(line string) position {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return position{
		x: x,
		y: y,
	}
}

func parseInstruction(line string) instruction {
	line = strings.TrimPrefix(line, "fold along ")
	parts := strings.Split(line, "=")
	y, _ := strconv.Atoi(parts[1])

	return instruction{
		axis:  rune(parts[0][0]),
		index: y,
	}
}

package _2021

func GetInfraredThermalImagingCameraSystemCode(rawInput string) string {
	dots, foldInstructions := parse(rawInput)
	paper := newPaper(dots)
	for _, foldInstruction := range foldInstructions {
		paper.fold(foldInstruction)
	}
	return paper.print()
}

func (p *paper) print() string {
	s := ""
	for y := 0; y <= p.height; y++ {
		for x := 0; x <= p.width; x++ {
			dot := p.dots[position{
				x: x,
				y: y,
			}]
			c := '.'
			if dot {
				c = '#'
			}
			s += string(c)
		}
		s += "\n"
	}
	return s
}

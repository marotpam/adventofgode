package _2017

type point struct {
	x, y int
}

func (p *point) hammingDistanceTo(other point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

type matrix struct {
	layers, size int
}

func NewMatrixFitting(n int) *matrix {
	if n == 0 {
		return &matrix{}
	}

	layers, size := 1, 1
	for ; size*size < n; {
		size += 2
		layers++
	}

	return &matrix{
		layers: layers,
		size:   size,
	}
}

func (m *matrix) findPointOne() point {
	middle := m.size / 2

	return point{
		x: middle,
		y: middle,
	}
}

func (m *matrix) findPointInOuterLayer(n int) point {
	bottomRight := m.size * m.size
	bottomLeft := bottomRight - m.size + 1
	topLeft := bottomLeft - m.size + 1
	topRight := topLeft - m.size + 1

	if n >= bottomLeft {
		return point{
			x: m.size - 1,
			y: n - bottomLeft,
		}
	}

	if n >= topLeft {
		return point{
			x: n - topLeft,
			y: 0,
		}
	}

	if n >= topRight {
		return point{
			x: 0,
			y: (m.size - 1) - (n - topRight),
		}
	}

	return point{
		x: m.size - 1,
		y: topRight - n,
	}
}

func CalculateDistance(n int) int {
	m := NewMatrixFitting(n)
	o := m.findPointOne()

	return o.hammingDistanceTo(m.findPointInOuterLayer(n))
}

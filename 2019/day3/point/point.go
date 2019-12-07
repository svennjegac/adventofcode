package point

func New(x, y int) Point {
	return Point{x: x, y: y}
}

type Point struct {
	x int
	y int
}

func (p Point) Add(p2 Point) Point {
	return Point{x: p.x + p2.x, y: p.y + p2.y}
}

func (p Point) Distance() int {
	x := p.x
	if x < 0 {
		x = -x
	}
	y := p.y
	if y < 0 {
		y = -y
	}
	return x + y
}

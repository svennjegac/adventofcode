package point

type Point struct {
	X int
	Y int
}

func (p Point) Neighs() []Point {
	return []Point{
		{
			X: p.X - 1,
			Y: p.Y,
		},
		{
			X: p.X,
			Y: p.Y - 1,
		},
		{
			X: p.X + 1,
			Y: p.Y,
		},
		{
			X: p.X,
			Y: p.Y + 1,
		},
	}
}

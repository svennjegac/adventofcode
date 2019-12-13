package moon

import (
	"fmt"
)

type Moon struct {
	Id       int
	Position *Point
	Velocity *Point
}

func (m *Moon) TotalEnergy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func (m *Moon) PotentialEnergy() int {
	return m.Position.Abs()
}

func (m *Moon) KineticEnergy() int {
	return m.Velocity.Abs()
}

func (m *Moon) Move() {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y
	m.Position.Z += m.Velocity.Z
}

func (m *Moon) Copy() *Moon {
	return &Moon{
		Id:       m.Id,
		Position: m.Position.Copy(),
		Velocity: m.Velocity.Copy(),
	}
}

func (m *Moon) Equals(other *Moon) bool {
	return m.Id == other.Id && m.Position.Equals(other.Position) && m.Velocity.Equals(other.Velocity)
}

func (m *Moon) String() string {
	return fmt.Sprintf("Id:%d,Pos:%s,Vel:%s", m.Id, m.Position.String(), m.Velocity.String())
}

type Point struct {
	X, Y, Z int
}

func (p *Point) Abs() int {
	return abs(p.X) + abs(p.Y) + abs(p.Z)
}

func (p *Point) XC() int {
	return p.X
}

func (p *Point) YC() int {
	return p.Y
}

func (p *Point) ZC() int {
	return p.Z
}

func (p *Point) Copy() *Point {
	return &Point{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p *Point) String() string {
	return fmt.Sprintf("x:%d,y:%d,z:%d", p.X, p.Y, p.Z)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

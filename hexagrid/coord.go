package hexagrid

import (
	"errors"
	"math"
)

// Coord is a hexagrid coordinate
type Coord struct {
	q int
	r int
	s int
}

// NewCoord returns a new Hexagrid Coord
func NewCoord(q, r, s int) (coord Coord, err error) {
	if (q + s + r) != 0 {
		err = errors.New("invalid coordinate value, q+r+s must be 0")
		return
	}
	coord = Coord{q, r, s}
	return
}

// Distance returns the distance between two Coords as a integer
func Distance(c1, c2 Coord) int {
	vec := Subtract(c1, c2)
	return int(math.Abs(float64(vec.q))+math.Abs(float64(vec.r))+math.Abs(float64(vec.s))) / 2
}

// Same check if coords are the same
func Same(c1, c2 Coord) bool {
	return c1.q == c2.q && c1.r == c2.r && c1.s == c2.s
}

// Neighboors of a given Coord
func (c Coord) Neighboors() []Coord {
	result := []Coord{}
	for _, v := range Neighboors {
		result = append(result, Add(c, v))
	}
	return result
}

// Add two Coords
func Add(c1, c2 Coord) Coord {
	return Coord{
		q: c1.q + c2.q,
		r: c1.r + c2.r,
		s: c1.s + c2.s,
	}
}

// Subtract two Coords
func Subtract(a, b Coord) Coord {
	return Coord{
		q: a.q - b.q,
		r: a.r - b.r,
		s: a.s - b.s,
	}
}

// Mult multiplies `m` to a Coord
func Mult(a Coord, m int) Coord {
	return Coord{
		q: a.q * m,
		r: a.r * m,
		s: a.s * m,
	}
}

// ToAxial converts Coord to Axial Coords
func (c Coord) ToAxial() Axial {
	return NewAxial(c.q, c.r)
}

// Axial Coords
type Axial struct {
	q int
	r int
}

// NewAxial creates a new Axial Coord
func NewAxial(q, r int) Axial {
	return Axial{q, r}
}

// ToCoord converts Axial coord to Coord (3 coords)
func (a Axial) ToCoord() (c Coord) {
	c, _ = NewCoord(a.q, a.r, -a.q-a.r)
	return
}

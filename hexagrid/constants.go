package hexagrid

// Coord Directions for flat hexagrid
var (
	NorthEast = Coord{1, -1, 0}
	SouthEast = Coord{1, 0, -1}
	South     = Coord{0, 1, -1}
	SouthWest = Coord{-1, 1, 0}
	NorthWest = Coord{-1, 0, 1}
	North     = Coord{0, -1, 1}
)

// Neighboors for flat hexagrid
var (
	Neighboors = []Coord{
		NorthEast,
		SouthEast,
		South,
		SouthWest,
		NorthWest,
		North,
	}
)

// Diagonal Directions for flat hexagrid
var (
	TopEast    = Coord{1, -2, 1}
	East       = Coord{2, -1, -1}
	BottomEast = Coord{1, 1, -2}
	BottomWest = Coord{-1, 2, -1}
	West       = Coord{-2, 1, 1}
	TopWest    = Coord{-1, -1, 2}
)

// Diagonals for flat hexagrid
var (
	Diagonals = []Coord{
		TopEast,
		East,
		BottomEast,
		BottomWest,
		West,
		TopWest,
	}
)

package hexagrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Coord(t *testing.T) {

	coord, err := NewCoord(1, -1, 0)
	if err != nil {
		fmt.Println("Coordinate unexpected error ", err)
		t.Fail()
	}
	if coord.q != 1 || coord.r != -1 || coord.s != 0 {
		fmt.Println("Error creating coordinate.\nexpected: ", 1, -1, 0, "\nactual: ", coord.q, coord.r, coord.s)
		t.Fail()
	}

	coord, err = NewCoord(1, 0, 0)
	if err == nil {
		fmt.Println("Expected error got nil")
		t.Fail()
	}

	coord = NewAxial(1, 0).ToCoord()
	if coord.q != 1 || coord.r != 0 || coord.s != -1 {
		fmt.Println("Error creating coordinate.\nexpected: ", 1, 0, -1, "\nactual: ", coord.q, coord.r, coord.s)
		t.Fail()
	}

}

func Test_Distance(t *testing.T) {
	_, err := NewCoord(1, -1, 0)
	if err != nil {
		fmt.Println("Coordinate unexpected error ", err)
		t.Fail()
	}
	c2, err := NewCoord(0, 0, 0)
	if err != nil {
		fmt.Println("Coordinate unexpected error ", err)
		t.Fail()
	}

	for _, n := range Neighboors {
		distance := Distance(c2, n)
		assert.Equal(t, 1, distance, "Distance calcs")
		distance = Distance(c2, Mult(n, 2))
		assert.Equal(t, 2, distance, "Distance calcs")
	}

}

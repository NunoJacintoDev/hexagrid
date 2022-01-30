package hexagrid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HexaGrid(t *testing.T) {

	t.Run("grid_1", func(t *testing.T) {
		grid := loadGrid(t, "../test/data/grid_1.json")

		start, err := NewCoord(0, 0, 0)
		assert.NoError(t, err)

		goal, err := NewCoord(-1, 1, 0)
		assert.NoError(t, err)

		t.Run("search_path", func(t *testing.T) {
			result, err := grid.SearchPath(start, goal, SearchPathOptions{})
			assert.NoError(t, err)
			assert.Equal(t, result, []Coord{start, goal})
		})

		t.Run("invalid_start", func(t *testing.T) {
			start, err = NewCoord(2, -2, 0)
			assert.NoError(t, err)
			_, err := grid.SearchPath(start, goal, SearchPathOptions{})
			assert.Error(t, err)
		})

		t.Run("invalid_goal", func(t *testing.T) {
			goal, err = NewCoord(2, -2, 0)
			assert.NoError(t, err)
			_, err := grid.SearchPath(start, goal, SearchPathOptions{})
			assert.Error(t, err)
		})
	})

	t.Run("grid_2", func(t *testing.T) {
		grid := loadGrid(t, "../test/data/grid_2.json")

		start, err := NewCoord(0, 0, 0)
		assert.NoError(t, err)

		t.Run("search_path", func(t *testing.T) {
			for _, goal := range Neighboors {
				result, err := grid.SearchPath(start, goal, SearchPathOptions{})
				assert.NoError(t, err)
				assert.Equal(t, []Coord{start, goal}, result)

				goalTwice := Mult(goal, 2)
				result, err = grid.SearchPath(start, goalTwice, SearchPathOptions{})
				assert.NoError(t, err)
				assert.Equal(t, []Coord{start, goal, goalTwice}, result)
			}
		})

		t.Run("unreachable_goal", func(t *testing.T) {
			goal := Coord{-5, 5, 0}
			_, err := grid.SearchPath(start, goal, SearchPathOptions{})
			assert.Error(t, err)
		})

	})

	t.Run("grid_3", func(t *testing.T) {
		grid := loadGrid(t, "../test/data/grid_3.json")

		start, err := NewCoord(0, 0, 0)
		assert.NoError(t, err)

		goal, err := NewCoord(-2, 2, 0)
		assert.NoError(t, err)

		t.Run("search_path", func(t *testing.T) {
			result, err := grid.SearchPath(start, goal, SearchPathOptions{})
			assert.NoError(t, err)
			assert.Equal(t, []Coord{start, {-1, 0, 1}, {-2, 1, 1}, goal}, result)
		})
	})

	// Add a restriction from one to other
	t.Run("grid_4", func(t *testing.T) {
		grid := loadGrid(t, "../test/data/grid_4.json")

		start, err := NewCoord(0, 0, 0)
		assert.NoError(t, err)

		goal := Coord{-1, 0, 1}

		t.Run("search_path_with_steps", func(t *testing.T) {
			result, err := grid.SearchPath(start, goal, SearchPathOptions{ApplyStep: true, MaxStep: 2})
			assert.NoError(t, err)
			assert.Equal(t, []Coord{{0, 0, 0}, {-1, 1, 0}, {0, 1, -1}, {1, 0, -1}, {1, -1, 0}, {0, -1, 1}, {-1, 0, 1}}, result)
		})
	})

	t.Run("get_set_value", func(t *testing.T) {
		h := NewHexagrid()
		h.Add(Coord{0, 0, 0}, 2, 0)
		value, exists := h.Get(Coord{0, 0, 0})
		assert.Equal(t, 2, value)
		assert.Equal(t, true, exists)

		value, exists = h.Get(Coord{0, 1, -1})
		assert.Equal(t, nil, value)
		assert.Equal(t, false, exists)

	})

}

func loadGrid(t *testing.T, path string) (h Hexagrid) {

	h = NewHexagrid()

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var grid dummyTestGrid

	json.Unmarshal(byteValue, &grid)

	for _, p := range grid.Pieces {
		c, err := NewCoord(p.Q, p.R, p.S)
		if err != nil {
			t.Fatal("Failed building Grid", err)
		}
		h.Add(c, "", float64(p.Cost))
	}

	return
}

type dummyTestGrid struct {
	Pieces []dummyTestPiece `json:"pieces"`
}

type dummyTestPiece struct {
	Q    int `json:"q"`
	R    int `json:"r"`
	S    int `json:"s"`
	Cost int `json:"cost"`
}

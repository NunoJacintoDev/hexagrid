package hexagrid

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
)

// Hexagrid struct with pieces and a movement cost
type Hexagrid struct {
	pieces map[Coord]interface{}
	cost   map[Coord]float64
}

// NewHexagrid creates a new Hexagrid
func NewHexagrid() Hexagrid {
	pieces := make(map[Coord]interface{})
	cost := make(map[Coord]float64)
	return Hexagrid{pieces: pieces, cost: cost}
}

// Add a new Piece to the Grid at `Coord` with a given movement `Cost` or altitude
func (h *Hexagrid) Add(c Coord, value interface{}, cost float64) {
	h.pieces[c] = value
	h.cost[c] = cost
}

// Remove a Piece from the Grid
func (h *Hexagrid) Remove(c Coord) {
	delete(h.pieces, c)
	delete(h.cost, c)
}

// Get Hexagrid Piece value at `Coord`
func (h *Hexagrid) Get(c Coord) (value interface{}, ok bool) {
	value, ok = h.pieces[c]
	return
}

// Cost Get cost/altitude of Hexagrid Piece value at `Coord`
func (h *Hexagrid) Cost(c Coord) float64 {
	return h.cost[c]
}

// Has check if Grid has a Piece
func (h *Hexagrid) Has(c Coord) bool {
	_, ok := h.pieces[c]
	return ok
}

// Neighboors of a Coord in the Hexagrid
func (h *Hexagrid) Neighboors(c Coord) []Coord {
	result := []Coord{}
	for _, v := range Neighboors {
		n := Add(c, v)
		if _, ok := h.pieces[n]; ok {
			result = append(result, n)
		}
	}
	return result
}

// ReachableNeighboors of a Coord in the Hexagrid for a given cost/altitude diff
func (h *Hexagrid) ReachableNeighboors(c Coord, diff float64) []Coord {
	result := []Coord{}
	for _, v := range Neighboors {
		n := Add(c, v)
		if _, ok := h.pieces[n]; ok {
			if _, ok := h.cost[n]; ok && math.Abs(h.costDiff(c, n)) <= diff {
				result = append(result, n)
			}
		}
	}
	return result
}

type SearchPathOptions struct {
	MaxStep   float64
	ApplyStep bool
}

// SearchPath, get shortest path from a start to a goal position in the hexagrid (uses A* Search Algorithm)
func (h *Hexagrid) SearchPath(start, goal Coord, options SearchPathOptions) (result []Coord, err error) {

	if !h.Has(start) {
		err = errors.New(fmt.Sprint("Piece not found", start))
		return
	}
	if !h.Has(goal) {
		err = errors.New(fmt.Sprint("Piece not found", goal))
		return
	}

	frontier := make(PriorityQueue, 1)

	startItem := &Item{
		value:    start,
		priority: 0,
		index:    0,
	}
	frontier[0] = startItem
	heap.Init(&frontier)

	cameFrom := make(map[Coord]interface{})
	costSoFar := make(map[Coord]float64)

	cameFrom[start] = nil
	costSoFar[start] = 0

	for len(frontier) != 0 {
		currentItem := heap.Pop(&frontier).(*Item)
		current := (currentItem.value).(Coord)

		if Same(current, goal) {
			return buildPath(cameFrom, start, goal)
		}

		var neighboors []Coord

		// apply step diff
		if options.ApplyStep {
			neighboors = h.ReachableNeighboors(current, options.MaxStep)
		} else {
			neighboors = h.Neighboors(current)
		}

		for _, next := range neighboors {
			newCost := costSoFar[current] + h.costDiff(current, next)

			if value, ok := costSoFar[next]; !ok || newCost < value {
				costSoFar[next] = newCost
				priority := -(newCost + heuristic(goal, next))
				newItem := &Item{
					value:    next,
					priority: int(priority),
				}
				heap.Push(&frontier, newItem)
				frontier.update(newItem, newItem.value, int(priority))

				cameFrom[next] = current
			}
		}
	}
	err = errors.New(fmt.Sprint("unreachable position", goal))

	return
}

// buildPath auxiliar function for path finding
func buildPath(cameFrom map[Coord]interface{}, start Coord, goal Coord) (result []Coord, err error) {
	current := goal
	from, exist := cameFrom[current]
	if !exist {
		err = errors.New(fmt.Sprint("path not found from ", start, "to", goal))
		return
	}
	result = append(result, current)

	for exist && from != nil {
		result = append(result, from.(Coord))

		current = from.(Coord)
		from, exist = cameFrom[current]
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return
}

// costDiff calculate cost diff between two coords
func (h *Hexagrid) costDiff(c, other Coord) float64 {
	return h.cost[other] - h.cost[c]
}

// heuristic function to path finding
func heuristic(a, b Coord) float64 {
	return math.Abs(float64(a.q-b.q)) + math.Abs(float64(a.r-b.r)) + math.Abs(float64(a.s-b.s))
}

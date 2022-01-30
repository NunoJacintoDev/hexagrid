package hexagrid

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Priority_Queue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, item.value, "orange")
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, item.value, "pear")
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, item.value, "banana")
	item = heap.Pop(&pq).(*Item)
	assert.Equal(t, item.value, "apple")

}

package main

import "fmt"

// node of a circular, doubled linked list
type marble struct {
	num  int
	next *marble
	prev *marble
}

func newMarbleCircle(num int) *marble {
	m := marble{}
	m.num = num
	m.next = &m
	m.prev = &m

	return &m
}

// inserts after m, returns the new marble
func (m *marble) insertAfter(newValue int) *marble {
	new := marble{}
	new.num = newValue
	new.next = m.next
	new.prev = m

	m.next.prev = &new
	m.next = &new

	return &new
}

// removes m, returns its value
func (m *marble) remove() int {
	m.prev.next = m.next
	m.next.prev = m.prev
	return m.num
}

// these come from input
const (
	numPlayers = 400
	lastMarble = 71864 * 100
)

func main() {
	scores := make([]int, numPlayers)
	currentMarble := newMarbleCircle(0) // insert first marble

	for i := 1; i <= lastMarble; i++ {
		currentPlayer := (i - 1) % numPlayers

		if i%23 != 0 {
			currentMarble = currentMarble.next.insertAfter(i)
		} else {
			scores[currentPlayer] += i
			toRemove := currentMarble.prev.prev.prev.prev.prev.prev.prev // not exactly elegant :-)
			currentMarble = toRemove.next
			scores[currentPlayer] += toRemove.remove()
		}
	}

	// find highest score
	highest := 0
	for _, v := range scores {
		if v > highest {
			highest = v
		}
	}

	fmt.Printf("Highest score is %v\n", highest)
}

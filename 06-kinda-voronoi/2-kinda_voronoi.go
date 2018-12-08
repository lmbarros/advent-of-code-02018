package main

import (
	"fmt"
)

const size = 400
const numLocs = 50

var centers [numLocs + 1]coord // index 0 is not used

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type coord struct {
	x int
	y int
}

func dist(c1, c2 coord) int {
	return abs(c2.x-c1.x) + abs(c2.y-c1.y)
}

func distancesSum(c coord) int {
	result := 0

	for i := 1; i <= numLocs; i++ {
		result += dist(c, centers[i])
	}

	return result
}

func main() {
	// Read input
	var err error
	id := 0
	for err == nil {
		id++
		var c coord
		_, err = fmt.Scanf("%v, %v\n", &c.x, &c.y)
		if err != nil {
			continue
		}

		centers[id] = c
	}

	// Find nearest to each cell
	safeArea := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if distancesSum(coord{i, j}) < 10000 {
				safeArea++
			}
		}
	}

	fmt.Printf("Safe area: %v\n", safeArea)
}

package main

import (
	"fmt"
)

const size = 400
const numLocs = 50

var m [400][400]int              // the nearest location ID
var isInfinite [numLocs + 1]bool // index 0 is not used
var areas [numLocs + 1]int       // index 0 is not used
var centers [numLocs + 1]coord   // index 0 is not used

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

func findNearest(c coord) int {
	nearest := 0
	minDist := 9999999
	tied := false

	for i := 1; i <= numLocs; i++ {
		d := dist(c, centers[i])
		if d < minDist {
			tied = false
			minDist = d
			nearest = i
		} else if d == minDist {
			tied = true
		}
	}

	if tied {
		return 0
	}

	if c.x == 0 || c.x == size-1 || c.y == 0 || c.y == size-1 {
		isInfinite[nearest] = true
	}

	return nearest
}

func dumpM() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(m[i][j])
			if j != size-1 {
				fmt.Print(",")
			}
		}
		fmt.Println()
	}
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
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			nearest := findNearest(coord{i, j})
			m[i][j] = nearest //not really needed, may be good for debugging
			areas[nearest]++  // yeah, even at [0]; let's count "borders", too
		}
	}

	// Find the largest area
	largest := 0
	for i := 1; i <= numLocs; i++ {
		if areas[i] > largest && !isInfinite[i] {
			largest = areas[i]
		}
	}

	// fmt.Printf("Areas: %v\n", areas)
	fmt.Printf("Largest: %v\n", largest)
	//	dumpM()
}

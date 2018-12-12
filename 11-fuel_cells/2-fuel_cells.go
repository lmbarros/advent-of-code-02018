package main

import "fmt"

var grid [300][300]int

const serialNumber = 8868

func getHundreds(v int) int {
	if v < 100 {
		return 0
	}
	v -= v / 1000 * 1000

	return v / 100
}

func main() {

	// Fill cell values
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			x := i + 1
			y := j + 1

			rackID := x + 10
			power := rackID * y
			power += serialNumber
			power *= rackID
			power = getHundreds(power)
			power -= 5

			grid[i][j] = power
		}
	}

	// Find the top nxn cell, without even trying to be smart. Let's just hope
	// for the best.
	largestPower := -9999999
	lpX := -1
	lpY := -1
	lpSize := -1

	for s := 1; s <= 300; s++ { // cell size
		for i := 0; i <= 300-s; i++ { // top-...
			for j := 0; j <= 300-s; j++ { // ...left corner
				power := 0
				for ci := i; ci < i+s; ci++ { // cell i
					for cj := j; cj < j+s; cj++ { // cell j
						power += grid[ci][cj]
					}
				}
				if power > largestPower {
					largestPower = power
					lpX = i + 1
					lpY = j + 1
					lpSize = s
				}
			}
		}
	}

	// There we are
	fmt.Printf("Largest power is %v,%v,%v\n", lpX, lpY, lpSize)
}

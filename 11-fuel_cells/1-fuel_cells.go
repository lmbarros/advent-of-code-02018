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

	// Find the top 3x3 cell, without even trying to be smart
	largestPower := -9999999
	lpX := -1
	lpY := -1
	for i := 0; i <= 297; i++ {
		for j := 0; j <= 297; j++ {
			power := grid[i][j] + grid[i+1][j] + grid[i+2][j] +
				grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1] +
				grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]

			if power > largestPower {
				largestPower = power
				lpX = i + 1
				lpY = j + 1
			}
		}
	}

	// There we are
	fmt.Printf("Largest power is at %v,%v\n", lpX, lpY)
}

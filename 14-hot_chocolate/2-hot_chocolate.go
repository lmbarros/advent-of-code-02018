package main

import (
	"fmt"
	"os"
)

var input = []int{2, 8, 6, 0, 5, 1}

var recipes []int
var current [2]int

func main() {
	// Init stuff
	recipes = make([]int, 2, 100*1024)
	recipes[0] = 3
	recipes[1] = 7
	current[0] = 0
	current[1] = 1

	for true {
		// Combine current recipes, add to recipes list
		sum := recipes[current[0]] + recipes[current[1]]
		if sum < 10 {
			recipes = append(recipes, sum)
		} else {
			recipes = append(recipes, sum/10) // always 1, right?
			recipes = append(recipes, sum%10)
		}

		// Advance
		for i, v := range current {
			v += 1 + recipes[current[i]]
			current[i] = v % len(recipes)
		}

		// Check match
		if len(recipes) >= len(input) {
			match := true

			j := len(input) - 1
			for i := len(recipes) - 1; i >= len(recipes)-len(input); i-- {
				if recipes[i] != input[j] {
					match = false
					break
				}
				j--
			}

			if match {
				res := len(recipes) - len(input)
				fmt.Printf("Recipes before input = %v\n", res)
				os.Exit(0)
			}
		}

		// Check match again, one position before. Ugly, but solves the
		// unforseen case arising from the fact that we might grow the recipe
		// list by either one or two each round.
		if len(recipes) >= len(input) {
			match := true

			j := len(input) - 1
			for i := len(recipes) - 2; i >= len(recipes)-len(input)-1; i-- {
				if recipes[i] != input[j] {
					match = false
					break
				}
				j--
			}

			if match {
				res := len(recipes) - len(input) - 1
				fmt.Printf("Recipes before input = %v\n", res)
				os.Exit(0)
			}
		}
	}
}

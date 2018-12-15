package main

import "fmt"

const recipesToTry = 286051 // the input

var recipes []int
var current [2]int

func main() {
	// Init stuff
	recipes = make([]int, 2, recipesToTry+10)
	recipes[0] = 3
	recipes[1] = 7
	current[0] = 0
	current[1] = 1

	for len(recipes) < recipesToTry+10 {
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
	}

	fmt.Print("Interesting recipes: ")
	for _, v := range recipes[recipesToTry : recipesToTry+10] {
		fmt.Printf("%v", v)
	}
	fmt.Println()
}

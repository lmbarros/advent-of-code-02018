package main

import "fmt"

const (
	numPots        = 200
	zero           = 50
	numGenerations = 20
)

var pots = [numPots]bool{}

var rules = make(map[[5]bool]bool)

func main() {
	// Read and set initial state
	var initialState string
	_, _ = fmt.Scanf("initial state: %v", &initialState)
	for i, v := range initialState {
		pots[zero+i] = v == '#'
	}

	// Skip that empty line
	_, _ = fmt.Scanf("\n")

	// Read and set rules
	var err error
	for err == nil {
		var lhs, rhs string
		_, err = fmt.Scanf("%v => %v", &lhs, &rhs)
		if err != nil {
			continue
		}

		lhsBool := [5]bool{lhs[0] == '#', lhs[1] == '#', lhs[2] == '#', lhs[3] == '#', lhs[4] == '#'}
		rhsBool := rhs[0] == '#'
		rules[lhsBool] = rhsBool
	}

	// Run the simulation
	for i := 0; i < numGenerations; i++ {
		nextPots := pots
		for j := 2; j <= numPots-3; j++ {
			ns := pots[j-2 : j+3] // neighbors slice
			neighbors := [5]bool{ns[0], ns[1], ns[2], ns[3], ns[4]}
			nextPots[j] = rules[neighbors]
		}
		pots = nextPots
	}

	// Compute final value
	result := 0
	for ai, v := range pots { // ai = absolute index
		if v {
			result += ai - zero
		}
	}

	fmt.Printf("Result is %v\n", result)
}

package main

import "fmt"

// returns the value of the node at the root
func readSubtreeValue(input *[]int) int {
	numChildren := (*input)[0]
	numMetadata := (*input)[1]
	*input = (*input)[2:]

	childValues := make([]int, 0, numChildren)
	for i := 0; i < numChildren; i++ {
		childValues = append(childValues, readSubtreeValue(input))
	}

	sum := 0

	if numChildren == 0 {
		// Value is sum of metadata
		for i := 0; i < numMetadata; i++ {
			sum += (*input)[i]
		}
	} else {
		// Value is value of children indexed by metadata
		for i := 0; i < numMetadata; i++ {
			index := (*input)[i]
			if index <= numChildren {
				sum += childValues[index-1]
			}
		}
	}

	*input = (*input)[numMetadata:]
	return sum
}

func main() {
	input := make([]int, 0, 20*1024)

	// Read input
	var err error
	for err == nil {
		var n int
		_, err = fmt.Scanf("%v", &n)
		if err != nil {
			continue
		}

		input = append(input, n)
	}

	fmt.Printf("Sum of metadata: %v\n", readSubtreeValue(&input))
}

package main

import "fmt"

// returns the sum of the metadata of the root node and child nodes
func readSubtree(input *[]int) int {
	numChildren := (*input)[0]
	numMetadata := (*input)[1]
	*input = (*input)[2:]

	sum := 0
	for i := 0; i < numChildren; i++ {
		sum += readSubtree(input) // some recursion, at last :-)
	}

	for i := 0; i < numMetadata; i++ {
		sum += (*input)[i]
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

	fmt.Printf("Sum of metadata: %v\n", readSubtree(&input))
}

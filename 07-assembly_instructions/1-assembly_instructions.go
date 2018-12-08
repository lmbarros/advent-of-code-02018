package main

import (
	"fmt"
	"sort"
)

type step struct {
	name string
	deps []string
}

type steps []step

func (s steps) Len() int {
	return len(s)
}

func (s steps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s steps) Less(i, j int) bool {
	if len(s[i].deps) == len(s[j].deps) {
		return s[i].name < s[j].name
	}

	return len(s[i].deps) < len(s[j].deps)
}

var instructions steps

func removeDependency(name string) {
	for i := 0; i < len(instructions); i++ {
		for j := 0; j < len(instructions[i].deps); j++ {
			if instructions[i].deps[j] == name {
				instructions[i].deps = append(instructions[i].deps[:j], instructions[i].deps[j+1:]...)
				break
			}
		}
	}
	sort.Sort(instructions)
}

func main() {
	instructions = make(steps, 0, 1000)

	tempInput := map[string][]string{}

	// Read input
	var err error
	for err == nil {
		var name, dep string
		_, err = fmt.Scanf("Step %v must be finished before step %v can begin.\n", &dep, &name)
		if err != nil {
			continue
		}

		tempInput[name] = append(tempInput[name], dep)
		_, depExists := tempInput[dep]
		if !depExists {
			tempInput[dep] = []string{}
		}
	}

	// Copy input to definitive data structure
	for k, v := range tempInput {
		instructions = append(instructions, step{k, v})
	}
	sort.Sort(instructions)

	// Process steps
	for len(instructions) > 0 && len(instructions[0].deps) == 0 {
		s := instructions[0].name
		fmt.Print(s)
		instructions = instructions[1:]
		removeDependency(s)
	}
	fmt.Println()

	// Debugging
	//fmt.Println(tempInput)
	//fmt.Println(instructions)

}

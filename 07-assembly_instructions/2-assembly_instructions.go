package main

import (
	"fmt"
	"sort"
)

const numWorkers = 5

type step struct {
	name     string
	deps     []string
	duration int
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

type scheduledTask struct {
	name       string
	finishTime int
}

type taskSchedule []scheduledTask

func (ts taskSchedule) Len() int           { return len(ts) }
func (ts taskSchedule) Swap(i, j int)      { ts[i], ts[j] = ts[j], ts[i] }
func (ts taskSchedule) Less(i, j int) bool { return ts[i].finishTime < ts[j].finishTime }

var instructions steps
var sched taskSchedule

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
	sched = make([]scheduledTask, 0, 1000)
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
		instructions = append(instructions, step{k, v, 60 + int(k[0]) - 'A' + 1})
	}
	sort.Sort(instructions)

	// Process steps
	freeWorkers := numWorkers
	time := 0

	for true {
		// Schedule next batch of instructions
		for len(instructions) > 0 && len(instructions[0].deps) == 0 && freeWorkers > 0 {
			s := instructions[0]
			freeWorkers--
			sched = append(sched, scheduledTask{s.name, time + s.duration})
			instructions = instructions[1:]
		}

		// Wait until finishing next task (or tasks, if finishing time is the same)
		time += sched[0].finishTime - time
		for len(sched) > 0 && sched[0].finishTime == time {
			freeWorkers++
			removeDependency(sched[0].name)
			sched = sched[1:]
		}

		// Are we ready?
		if len(instructions) == 0 && len(sched) == 0 {
			break
		}
	}

	fmt.Printf("Total time = %v\n", time)
}

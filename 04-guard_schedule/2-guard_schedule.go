package main

import (
	"fmt"
	"strconv"
)

func main() {
	sleepMap := map[int][]int{} // guard id -> times asleep at each minute

	var err error
	currentGuard := -1
	sleepStart := -1

	// Read input, fill data structures
	for err == nil {
		var year, month, day, h, m int
		var e1, e2 string

		_, err = fmt.Scanf("[%d-%d-%d %d:%d] %v %v", &year, &month, &day, &h, &m, &e1, &e2)
		if err != nil {
			continue
		}

		// Ouch, I didn't expect having to do this until much later...
		if e1[0] == 'G' {
			var devnull string
			_, err = fmt.Scanf("%v %v\n", &devnull, &devnull)
		}

		switch e1[0] {
		// New guard
		case 'G':
			currentGuard, _ = strconv.Atoi(e2[1:])
			_, exists := sleepMap[currentGuard]
			if !exists {
				sleepMap[currentGuard] = make([]int, 60)
			}

		// Falls asleep
		case 'f':
			sleepStart = m

		// Wakes up
		case 'w':
			for i := sleepStart; i < m; i++ {
				sleepMap[currentGuard][i]++
			}
		}
	}

	// Find the guard who slept most
	mostMinutesSoFar := 0
	sleepiestMinuteSoFar := 0
	sleepiestGuardSoFar := -1

	for g, s := range sleepMap {
		for min, v := range s {
			if v > mostMinutesSoFar {
				mostMinutesSoFar = s[min]
				sleepiestMinuteSoFar = min
				sleepiestGuardSoFar = g
			}
		}
	}

	// Phew, we are done
	fmt.Printf("Aaaaand the answer is %v\n", sleepiestGuardSoFar*sleepiestMinuteSoFar)
}

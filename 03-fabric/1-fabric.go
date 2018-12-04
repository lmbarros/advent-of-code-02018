package main

import "fmt"

const N = 1000

func main() {
	var f [N][N]int // fabric

	var err error
	for err == nil {
		var id, x, y, w, h int
		_, err = fmt.Scanf("#%v @ %v,%v: %vx%v\n", &id, &x, &y, &w, &h)
		if err != nil {
			continue
		}

		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				f[i][j]++
			}
		}
	}

	overlaps := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if f[i][j] > 1 {
				overlaps++
			}
		}
	}

	fmt.Printf("Ho ho ho! (%v)\n", overlaps)
}

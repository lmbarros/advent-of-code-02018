package main

import "fmt"

const N = 1000

type Patch struct {
	id, x, y, w, h int
}

func main() {
	var f [N][N]int // fabric

	patches := []Patch{}

	var err error
	for err == nil {
		p := Patch{}
		_, err = fmt.Scanf("#%v @ %v,%v: %vx%v\n", &p.id, &p.x, &p.y, &p.w, &p.h)
		if err != nil {
			continue
		}

		for i := p.x; i < p.x+p.w; i++ {
			for j := p.y; j < p.y+p.h; j++ {
				f[i][j]++
			}
		}

		patches = append(patches, p)
	}

nextPatch:
	for _, p := range patches {
		for i := p.x; i < p.x+p.w; i++ {
			for j := p.y; j < p.y+p.h; j++ {
				if f[i][j] != 1 {
					continue nextPatch
				}
			}
		}
		fmt.Printf("Ho ho ho! (%v)\n", p.id)
		return
	}
}

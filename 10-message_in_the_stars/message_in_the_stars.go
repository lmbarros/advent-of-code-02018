package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/svgo"
)

type star struct {
	px int64
	py int64
	vx int64
	vy int64
}

// Updates the sky and draws it; returns the new squared sky area
func updateAndDrawSky(sky []star, timeInSecsBeforeUpdate int) int64 {
	// Find bounds
	minX := sky[0].px
	maxX := sky[0].px
	minY := sky[0].py
	maxY := sky[0].py

	for _, s := range sky {
		if s.px < minX {
			minX = s.px
		}
		if s.px > maxX {
			maxX = s.px
		}
		if s.py < minY {
			minY = s.py
		}
		if s.py > maxY {
			maxY = s.py
		}
	}

	skyWidth := maxX - minX
	skyHeight := maxY - minY

	// Init file
	fileName := fmt.Sprintf("sky-%05v.svg", timeInSecsBeforeUpdate+1)
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Update and draw
	canvas := svg.New(file)
	canvas.Startview(int(skyWidth), int(skyHeight), int(minX), int(minY), int(skyWidth), int(skyHeight))

	for i, s := range sky {
		s.px += s.vx
		s.py += s.vy
		sky[i] = s

		r := int(skyWidth) / 500
		if r == 0 {
			r = 1
		}
		canvas.Circle(int(s.px), int(s.py), r)
	}

	canvas.End()

	// Return current (squared) sky area
	return skyWidth * skyHeight
}

func main() {

	sky := make([]star, 0, 500)

	var err error
	for err == nil {
		var s star
		_, err = fmt.Scanf("position=<%v, %v> velocity=<%v, %v>\n", &s.px, &s.py, &s.vx, &s.vy)
		if err != nil {
			continue
		}

		sky = append(sky, s)
	}

	prevArea := updateAndDrawSky(sky, 0)
	for i := 1; true; i++ {
		newArea := updateAndDrawSky(sky, i)
		if newArea > prevArea {
			break
		} else {
			prevArea = newArea
		}
	}
}

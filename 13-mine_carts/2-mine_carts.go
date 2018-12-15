package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var tracks []string

type cart struct {
	dir      rune
	x        int
	y        int
	nextTurn int // mod 3 -> 0: left, 1: straight, 2: right
	crashed  bool
}

type cartList []cart

func (c cartList) Len() int      { return len(c) }
func (c cartList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c cartList) Less(i, j int) bool {
	if !c[i].crashed && c[j].crashed {
		return true
	} else if c[i].crashed && !c[j].crashed {
		return false
	}
	if c[i].y == c[j].y {
		return c[i].x < c[j].x
	}
	return c[i].y < c[j].y
}

var carts cartList

func handleCrashAt(selfC cart, selfI int) cart {
	for i, c := range carts {
		if i == selfI {
			continue
		}

		if c.x == selfC.x && c.y == selfC.y {
			c.crashed = true
			selfC.crashed = true
			carts[i] = c
			carts[selfI] = selfC
			break
		}
	}
	return selfC
}

func (c cart) String() string {
	if c.crashed {
		return fmt.Sprintf("[CRASHED %c %v,%v]", c.dir, c.x, c.y)
	}
	return fmt.Sprintf("[%c %v,%v]", c.dir, c.x, c.y)
}

func toRight(dir rune) rune {
	switch dir {
	case '>':
		return 'v'
	case '<':
		return '^'
	case '^':
		return '>'
	case 'v':
		return '<'
	}

	panic("Bad dir for toRight")
}

func toLeft(dir rune) rune {
	switch dir {
	case '>':
		return '^'
	case '<':
		return 'v'
	case '^':
		return '<'
	case 'v':
		return '>'
	}

	panic("Bad dir for toLeft")
}

func main() {
	tracks = make([]string, 0, 150)
	carts = make([]cart, 0, 100)

	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		for x, c := range line {
			if c == '>' || c == '<' {
				carts = append(carts, cart{c, x, y, 0, false})
				line = line[:x] + "-" + line[x+1:]
			} else if c == '^' || c == 'v' {
				carts = append(carts, cart{c, x, y, 0, false})
				line = line[:x] + "|" + line[x+1:]
			}
		}
		tracks = append(tracks, line)
		y++
	}

	// Simulate!
	for {
		sort.Sort(carts)
		for i, c := range carts {
			if c.crashed {
				continue
			}
			switch c.dir {
			case '>':
				c.x++
			case '<':
				c.x--
			case '^':
				c.y--
			case 'v':
				c.y++
			}

			c = handleCrashAt(c, i)

			switch tracks[c.y][c.x] {
			case '/':
				switch c.dir {
				case '^':
					c.dir = '>'
				case 'v':
					c.dir = '<'
				case '>':
					c.dir = '^'
				case '<':
					c.dir = 'v'
				}
			case '\\':
				switch c.dir {
				case '^':
					c.dir = '<'
				case 'v':
					c.dir = '>'
				case '>':
					c.dir = 'v'
				case '<':
					c.dir = '^'
				}

			case '+':
				switch c.nextTurn % 3 {
				case 0:
					c.dir = toLeft(c.dir)
				case 2:
					c.dir = toRight(c.dir)
				}
				c.nextTurn++
			}

			carts[i] = c
		}

		// Remove crashed carts
		sort.Sort(carts)
		firstCrashed := -1
		for i, c := range carts {
			if c.crashed {
				firstCrashed = i
				break
			}
		}
		if firstCrashed >= 0 {
			carts = carts[:firstCrashed]
		}

		// There can be only one
		if len(carts) == 1 {
			fmt.Printf("Last cart at %v,%v\n", carts[0].x, carts[0].y)
			os.Exit(0)
		}
	}
}

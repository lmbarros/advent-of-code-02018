package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func isInvCase(a, b rune) bool {
	if unicode.IsUpper(a) && unicode.IsLower(b) {
		return a == unicode.ToUpper(b)
	} else if unicode.IsLower(a) && unicode.IsUpper(b) {
		return a == unicode.ToLower(b)
	} else {
		return false
	}
}

func main() {
	stack := make([]byte, 0, 100*1024) // first time feeling a bit smart :-)
	input, _ := ioutil.ReadFile("input.txt")

	for _, c := range input {
		stack = append(stack, c)
		n := len(stack)
		if n >= 2 && isInvCase(rune(stack[n-1]), rune(stack[n-2])) {
			stack = stack[:n-2]
		}
	}

	fmt.Printf("Final polymer size: %v\n", len(stack))
}

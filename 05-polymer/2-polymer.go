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

func calcLengthRemoving(input []byte, r byte) int {
	stack := make([]byte, 0, 100*1024)
	for _, c := range input {
		if unicode.ToLower(rune(c)) == rune(r) {
			continue
		}
		stack = append(stack, c)
		n := len(stack)
		if n >= 2 && isInvCase(rune(stack[n-1]), rune(stack[n-2])) {
			stack = stack[:n-2]
		}
	}

	return len(stack)
}

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	bestLen := maxInt
	for c := 'a'; c <= 'z'; c++ {
		length := calcLengthRemoving(input, byte(c))

		if length < bestLen {
			bestLen = length
		}
	}

	fmt.Printf("Best polymer size: %v\n", bestLen)
}

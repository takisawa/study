package main

import (
	"fmt"
)

const intSize = 32 << (^uint(0) >> 63)

func fastAbs(x int) int {
	// m := -1 if x < 0. m := 0 otherwise.
	m := x >> (intSize - 1)

	// In two's complement representation, the negative number
	// of any number (except the smallest one) can be computed
	// by flipping all the bits and add 1. This is faster than
	// code with a branch.
	// See Hacker's Delight, section 2-4.
	return (x ^ m) - m
}

func slowAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(fastAbs(-1))
	fmt.Println(slowAbs(-1))
}

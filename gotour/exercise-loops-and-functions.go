package main

import (
	"fmt"
	"math"
)

const Offset = 0.00001

func Sqrt(x float64) (float64, int) {
	z := 1.0
	prev_z := float64(0)
	loop_num := 0

	for math.Abs(prev_z-z) > Offset {
		prev_z = z
		z -= (z*z - x) / (2 * z)
		loop_num += 1
	}

	return z, loop_num
}

func main() {
	for x := 0; x < 15; x++ {
		z, loop_num := Sqrt(float64(x))
		fmt.Println(fmt.Sprintf("Sqrt(%d)=", x), z)
		fmt.Println("loop_num=", loop_num)

		if loop_num > 10 {
			fmt.Println("over 10 loop")
		} else {
			fmt.Println("less than 10 loop")
		}
	}
}

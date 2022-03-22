package main

import (
	"fmt"
	"math"
)

const Small = 0.0000000001

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		dist := z*z - x
		z -= dist / (2 * z)
		if (dist < 0 && dist > -Small) || (dist >= 0 && dist < Small) {
			return z
		}
	}
}

func main() {
	x := 6.0
	fmt.Println("X =", x)
	fmt.Println("My   Sqrt:", Sqrt(x))
	fmt.Println("Math.Sqrt:", math.Sqrt(x))
}

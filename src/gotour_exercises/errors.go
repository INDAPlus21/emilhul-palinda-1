package main

import (
	"fmt"
)

const Small = 0.0000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := x / 2
		for {
			dist := z*z - x
			z -= dist / (2 * z)
			if (dist < 0 && dist > -Small) || (dist >= 0 && dist < Small) {
				return z, nil
			}
		}
	} else {
		return x, ErrNegativeSqrt(-2)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns the
// an int representing the i:th number
// in the fibonacci sequence.
func fibonacci() func(int) int {
	n_0, n_1 := 0, 1
	return func(i int) int {
		switch i {
		case 0:
			return 0
		case 1:
			return 1
		default:
			n := n_0 + n_1
			n_0 = n_1
			n_1 = n
			return n
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

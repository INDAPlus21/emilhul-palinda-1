package twopartsum

// sum the numbers in a and send the result on res.
func sum(a []int, res chan int) {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	sum := <-ch + <-ch
	return sum
}

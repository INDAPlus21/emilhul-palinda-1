package twopartsum

import "testing"

type testPair struct {
	array []int
	sum   int
}

var tests = []testPair{
	// Positive
	{[]int{1, 3, 5, 7, 2, 4, 8}, 30}, // even ammount
	{[]int{2, 3, 5, 10, 15, 5}, 40},  // odd ammount
	{[]int{100, 300, 250, 150}, 800},
	// Negative
	{[]int{-1, -3, -5, -7, -2, -4, -8}, -30},
	{[]int{-2, -3, -5, -10, -15, -5}, -40},
	{[]int{-100, -300, -250, -150}, -800},
	// Mixed
	{[]int{1, -3, -5, 7, -2, 4, 8}, 10},
	{[]int{2, 3, 5, -10, -15, 5}, -10},
	{[]int{-100, -300, 250, 150}, 0},
}

func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func TestPositive(t *testing.T) {
	for _, pair := range tests[:3] {
		sum := ConcurrentSum(pair.array)
		if sum != pair.sum {
			t.Error("For", pair.array, "expected", pair.sum, "got", sum)
		}
	}
}

func TestNegative(t *testing.T) {
	for _, pair := range tests[3:6] {
		sum := ConcurrentSum(pair.array)
		if sum != pair.sum {
			t.Error("For", pair.array, "expected", pair.sum, "got", sum)
		}
	}
}

func TestMixed(t *testing.T) {
	for _, pair := range tests[6:] {
		sum := ConcurrentSum(pair.array)
		if sum != pair.sum {
			t.Error("For", pair.array, "expected", pair.sum, "got", sum)
		}
	}
}

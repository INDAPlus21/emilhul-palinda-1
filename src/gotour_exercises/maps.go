package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, value := range words {
		count := m[value]
		m[value] = count + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

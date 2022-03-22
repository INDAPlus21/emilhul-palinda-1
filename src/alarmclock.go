package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		t := time.Now().Local().Format("The time is 15.04.05:")
		fmt.Printf("%s %s\n", t, text)
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 60*time.Second)
	select {}
}

package main

import (
	"fmt"
	"time"
)

// This is actually possible in Go!
func main() {
	fmt.Println("Creating 100,000 goroutines...")

	for i := 0; i < 100000; i++ {
		go func(id int) {
			time.Sleep(10 * time.Second)
			if id%10000 == 0 {
				fmt.Printf("Goroutine %d still running!\n", id)
			}
		}(i)
	}

	fmt.Println("All 100,000 goroutines started!")
	time.Sleep(15 * time.Second)
}

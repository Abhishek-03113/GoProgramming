package main

import (
	"fmt"
	"time"
)

func slowOperation() {
	fmt.Println("Starting slow operation...")
	time.Sleep(10 * time.Second) // Simulates slow work
	fmt.Println("Slow operation completed!")
}

func main() {
	go slowOperation()

	// Wait for 3 seconds, then give up
	time.Sleep(3 * time.Second)
	fmt.Println("Timeout! Moving on...")
}

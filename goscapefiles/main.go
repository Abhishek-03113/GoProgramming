package main

import (
	"context"
	"fmt"
	"time"
)

func nightShiftWorker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // Emergency stop button pressed!
			fmt.Printf("🚨 %s received shutdown signal - going home!\n", name)
			return
		default:
			fmt.Printf("👷 %s is working the night shift...\n", name)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Set up emergency shutdown after 3 seconds
	ctx, emergencyStop := context.WithTimeout(context.Background(), 1*time.Second)
	defer emergencyStop()

	go nightShiftWorker(ctx, "Alice")
	go nightShiftWorker(ctx, "Bob")

	time.Sleep(5 * time.Second)
	fmt.Println("🏢 Building is now closed")
}

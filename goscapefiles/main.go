package main

import (
	"fmt"
	"time"
)

func reportNews(reporter string, stories []string) {
	for i, story := range stories {
		fmt.Printf("📰 %s reporting: %s (Story %d)\n", reporter, story, i+1)
		time.Sleep(1 * time.Second) // Time to write article
	}
	fmt.Printf("✅ %s finished all stories!\n", reporter)
}

func main() {
	fmt.Println("📺 News Agency starts the day")

	// Different reporters covering different beats
	sportsStories := []string{"Football Match Results", "Basketball Finals", "Tennis Tournament"}
	techStories := []string{"New AI Release", "Smartphone Launch", "Tech IPO"}
	weatherStories := []string{"Tomorrow's Forecast", "Hurricane Update", "Seasonal Changes"}

	// All reporters work simultaneously
	go reportNews("Sports Reporter", sportsStories)
	go reportNews("Tech Reporter", techStories)
	go reportNews("Weather Reporter", weatherStories)

	// Wait for all reporters to finish
	time.Sleep(5 * time.Second)
	fmt.Println("📺 News Agency day ends")
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func brewCoffee(coffeeType string) {
//	fmt.Printf("☕ Starting to brew %s...\n", coffeeType)
//	time.Sleep(2 * time.Second) // Brewing time
//	fmt.Printf("✅ %s is ready!\n", coffeeType)
//}
//
//func main() {
//	fmt.Println("🏪 Coffee shop opens!")
//
//	// Without goroutines - customers wait in line
//	brewCoffee("Espresso")
//	brewCoffee("Latte")
//	brewCoffee("Cappuccino")
//
//	fmt.Println("\n--- Now with concurrent brewing ---\n")
//
//	// With goroutines - multiple coffee machines working
//	go brewCoffee("Espresso")   // Machine 1
//	go brewCoffee("Latte")      // Machine 2
//	go brewCoffee("Cappuccino") // Machine 3
//
//	time.Sleep(3 * time.Second) // Wait for all to complete
//	fmt.Println("🚪 Coffee shop closes!")
//}

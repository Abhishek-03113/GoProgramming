package main

import (
	greet "HandsOn1/Greetings"
	"fmt"
)

func main() {

	alice := "Alice"
	bob := "Bob"
	charlie := "Charlie"

	fmt.Println("Greeting for custom usernames ")

	fmt.Println(greet.GreetingMessage(alice))
	fmt.Println(greet.GreetingMessage(bob))
	fmt.Println(greet.GreetingMessage(charlie))

	fmt.Println()
	fmt.Println("Greetings for default name ")
	fmt.Println(greet.GreetingMessage(""))

}

package main

import "time"

func main() {
	emailMap := make(map[string]User)

	alice := RegisterUser("Alice", "Alice@gmail.com", []Activity{}, emailMap)
	bob := RegisterUser("Bob", "Bob@gmail.com", []Activity{}, emailMap)

	a := NewActivity("Texting ")

	alice.startActivity(a)
	alice.stopActivity(a)

	bob.startActivity(a)
	time.Sleep(2 * time.Second)
	bob.stopActivity(a)

	alice.PrintLog()
	bob.PrintLog()
}

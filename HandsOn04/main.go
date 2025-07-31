package main

import (
	"HandsOn04/Contact"
	"fmt"
)

func main() {

	alice := Contact.CreateContact(
		"Alice",
		"Alice@something.com",
		20,
		150,
		true,
	)

	bob := Contact.CreateContact(
		"Bob",
		"bob@gmail.com",
		25,
		180,
		false,
	)

	charlie := Contact.CreateContact(
		"Charlie",
		"charlie@gmail.com",
		30,
		175,
		true,
	)

	dave := Contact.CreateContact(
		"Dave",
		"dave@gmail.com",
		40,
		160,
		true,
	)

	emma := Contact.CreateContact(
		"Emma",
		"emma@something.com",
		22,
		165,
		true,
	)

	contacts := []Contact.Contact{alice, bob, charlie, dave, emma}

	fmt.Print("Contact List:\n")
	fmt.Printf("Total Contacts: %d\n", len(contacts))

	alice.GetProfileSummary()
	fmt.Println()
	alice.PrintProfile()
	fmt.Println("--------------------")

	bob.GetProfileSummary()
	fmt.Println()
	bob.PrintProfile()
	fmt.Println("--------------------")

	charlie.GetProfileSummary()
	fmt.Println()
	charlie.PrintProfile()
	fmt.Println("--------------------")

	dave.GetProfileSummary()
	fmt.Println()
	dave.PrintProfile()
	fmt.Println("--------------------")

	emma.GetProfileSummary()
	fmt.Println()
	emma.PrintProfile()
	fmt.Println("--------------------")
}

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
	for _, contact := range contacts {
		contact.GetProfileSummary()
		fmt.Print("\n")
		contact.PrintProfile()
		fmt.Print("\n")
		fmt.Print("----------------------\n")
	}
}

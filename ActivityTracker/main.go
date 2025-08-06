package main

import (
	"fmt"
	"time"
)

func main() {
	var Users []User

	alice, success := RegisterUser("Alice", "Alice@gmail.com", []*Activity{}, Users)
	if success {
		fmt.Printf("User %s created Successfully \n", alice.name)
	} else {
		fmt.Printf("User creation failed, email already exist \n")
	}

	A1 := Activity{
		action: "Watching Reels",
		start:  time.Now(),
		Time:   0,
		status: 0,
	}

	A2 := Activity{
		action: "Watching Youtube",
		start:  time.Now(),
		Time:   0,
		status: 0,
	}

	A3 := Activity{
		action: "Texting",
		start:  time.Now(),
		Time:   0,
		status: 0,
	}

	alice.startActivity(A1)
	time.Sleep(2 * time.Second) // Simulate some time passing
	alice.stopActivity(&A1)

	alice.startActivity(A2)
	time.Sleep(time.Second) // Simulate some time passing
	alice.stopActivity(&A2)

	alice.startActivity(A3)
	time.Sleep(time.Second) // Simulate some time passing
	alice.stopActivity(&A3)

	for _, act := range alice.activities {
		if act.status == Finished {
			fmt.Println(act.Log())
		} else {
			fmt.Println("Activity Still running or Activity Not Started Yet")
		}
	}
}

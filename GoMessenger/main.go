package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	name string
}

func (u User) sendMessage(msg string, msgChannel *chan string) {
	formattedMessage := fmt.Sprintf("%s : %s", u.name, msg)
	*msgChannel <- formattedMessage
}

func MessageService(username string, messageChannel <-chan string) {
	time.Sleep(1 * time.Second)

	for message := range messageChannel {
		fmt.Printf("\n%s Received a Message from %s \n", username, message)
	}

}

func main() {

	msgChannelAlice := make(chan string, 5)
	msgChannelBob := make(chan string, 5)

	alice := User{"Alice"}
	bob := User{"Bob"}

	go MessageService(bob.name, msgChannelBob)
	go MessageService(alice.name, msgChannelAlice)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var sender string
		var message string

		fmt.Printf("Who is sending (alice/bob/exit): ")
		scanner.Scan()
		sender = strings.ToLower(strings.TrimSpace(scanner.Text()))

		if strings.ToLower(sender) == "exit" {
			close(msgChannelAlice)
			close(msgChannelBob)
			break
		}

		fmt.Print("Enter message: ")
		scanner.Scan()
		message = scanner.Text()

		if strings.ToLower(sender) == "alice" {
			alice.sendMessage(message, &msgChannelBob)
		} else if strings.ToLower(sender) == "bob" {
			bob.sendMessage(message, &msgChannelAlice)
		} else {
			fmt.Println("Unknown sender.")
		}
	}
}

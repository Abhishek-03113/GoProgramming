package Greetings

const (
	name = "Json"
)

var greeting string = "Hello !! Nice to Meet You !! "

func GreetingMessage(username string) string {

	if len(username) > 0 {
		return greeting + username
	}

	return greeting + name
}

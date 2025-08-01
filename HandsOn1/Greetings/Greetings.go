package Greetings

const (
	name = "Json"
)

var greeting string = "Hello !! Nice to Meet You !! "

func GreetingMessage(usernames ...string) string {

	if len(usernames) > 0 {
		return greeting + usernames[0]
	}
	return greeting + name

}

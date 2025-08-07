package main

func main() {
	emailMap := make(map[string]User)

	alice := RegisterUser("Alice", "Alice@gmail.com", make(map[string]Activity), emailMap)
	bob := RegisterUser("Bob", "Bob@gmail.com", make(map[string]Activity), emailMap)

	a := NewActivity("Texting ")
	b := NewActivity("Watching reels")
	c := NewActivity("Playing games")

	d := NewActivity("Downloading Movie")
	alice.startActivity(d)
	alice.startActivity(a)
	//time.Sleep(time.Second)
	alice.stopActivity(a)

	alice.startActivity(b)
	//time.Sleep(time.Second)
	alice.stopActivity(b)

	bob.startActivity(a)
	//time.Sleep(time.Second)
	bob.stopActivity(a)

	bob.startActivity(c)
	//time.Sleep(time.Second)
	bob.stopActivity(c)

	bob.startActivity(b)
	//time.Sleep(time.Second)
	bob.stopActivity(b)

	alice.PrintLog()
	bob.PrintLog()
}

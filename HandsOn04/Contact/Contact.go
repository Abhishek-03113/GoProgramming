package Contact

import "fmt"

type Contact struct {
	name     string
	email    string
	age      int
	height   int
	isActive bool
}

func (c Contact) GetName() string {
	return c.name
}

func (c Contact) GetEmail() string {
	return c.email
}

func (c Contact) GetAge() int {
	return c.age
}

func (c Contact) GetHeight() int {
	return c.height
}

func (c Contact) IsActive() bool {
	return c.isActive
}

func (c Contact) GetProfileSummary() {
	fmt.Printf("Name: %s, Email: %s, Age: %d", c.name, c.email, c.age)
}

func (c Contact) PrintProfile() {
	fmt.Print("Profile of ", c.name, ":\n")

	fmt.Printf(" Name : %s\n Email: %s\n Age: %d\n Height: %d cm\n Active: %t\n", c.name, c.email, c.age, c.height, c.isActive)
}

func CreateContact(name, email string, age, height int, isActive bool) Contact {
	return Contact{
		name:     name,
		email:    email,
		age:      age,
		height:   height,
		isActive: isActive,
	}
}

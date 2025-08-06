package main

import "fmt"

func main() {
	value := 5
	p := &value

	fmt.Println(p)

	*p = 10

	var x *int

	fmt.Println(x == nil)

	fmt.Println(value)
}

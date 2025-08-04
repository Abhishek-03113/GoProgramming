package main

import "fmt"

func main() {
	var p *int

	a := 10
	p = &a

	fmt.Println("%d is stored at %p", a, p)
	fmt.Println(p)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 885, 5, 4, 2}

	sort.Slice(a, func(i, j int) bool {
		return a[j] < a[i]
	})

	fmt.Println(a)

	x := A{
		X: 10,
		Y: 20,
	}

	fmt.Println(x)

	About(1.000)
}

type A struct {
	X int
	Y float64
}

func (a A) String() string {
	return fmt.Sprintf("This is struct which has two variable X = %d and Y = %f", a.X, a.Y)
}

type B struct {
	X int
	Y float64
}

func About(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("this is a string")
	case float64:
		fmt.Println("this is a float")
	}
}

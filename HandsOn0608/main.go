package main

import "fmt"

func main() {

	slice := []int{
		1, 2, 3, 4, 5,
	}

	fmt.Println(slice)
	double(slice)
	fmt.Println(slice)

}

func double(slice []int) {
	for i := 0; i < len(slice); i++ {
		(slice)[i] *= 2
	}
}

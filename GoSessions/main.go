package main

import (
	"fmt"
)
import "sort"

func main() {
	scores := []int{1, 2, 3, 100, 5, 6, 7, 8, 9}
	chars := []string{"a", "a", "b", "c", "b"}
	var marks []int
	marks = append(marks, 1, 2, 3, 4, 5)

	for _, val := range marks {
		print(val)
	}
	sort.Ints(scores)
	fmt.Println(scores)

	var numbers []int // Slice (dynamic array)
	numbers = append(numbers, 1, 2, 3)
	fmt.Println(len(numbers)) // Length: 3
	fmt.Println(cap(numbers)) // Capacity: may be larger

	var charmap map[string]int
	charmap = make(map[string]int)

	for _, char := range chars {
		charmap[char]++
	}

	for key, value := range charmap {
		fmt.Println("Number", key, "occurs", value, "times")
	}
}

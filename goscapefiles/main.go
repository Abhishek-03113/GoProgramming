package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 885, 5, 4, 2}

	sort.Slice(a, func(i, j int) bool {
		return a[j] < a[i]
	})

	hi := ""

	fmt.Println(a)
}

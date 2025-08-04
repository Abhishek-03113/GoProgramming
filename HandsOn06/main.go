package main

import "fmt"

/*Write a function GroupByRemainder(nums []int, n int) map[int][]int that groups numbers by their remainder when divided by n.*/

func GroupByRemainders(nums []int, n int) map[int][]int {
	remainderMap := make(map[int][]int)

	for _, num := range nums {

		remainder := num % n
		remainderMap[remainder] = append(remainderMap[remainder], num)
	}

	return remainderMap
}

func main() {

	nums := []int{2, 8, 55, 10, 12, 5, 7, 3, 9}

	fmt.Println(GroupByRemainders(nums, 3))

}
